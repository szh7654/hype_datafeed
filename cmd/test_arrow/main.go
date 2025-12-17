package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/apache/arrow/go/v18/arrow"
	"github.com/apache/arrow/go/v18/arrow/array"
	"github.com/apache/arrow/go/v18/arrow/ipc"
	"github.com/apache/arrow/go/v18/arrow/memory"

	// 引入 Parquet 相关包
	"github.com/apache/arrow/go/v18/parquet"
	"github.com/apache/arrow/go/v18/parquet/compress"
	"github.com/apache/arrow/go/v18/parquet/pqarrow"
)

// 模拟网络 Writer
type MockNetworkWriter struct{}

func (m *MockNetworkWriter) Write(p []byte) (n int, err error) { return len(p), nil }

func main() {
	pool := memory.NewGoAllocator()

	// 定义 Schema
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "timestamp", Type: &arrow.TimestampType{Unit: arrow.Nanosecond}},
			{Name: "value", Type: arrow.PrimitiveTypes.Float64},
		}, nil,
	)

	// ---------------------------------------------------------
	// 1. 初始化网络流 (保持不变：网络继续用 Arrow IPC Stream)
	// ---------------------------------------------------------
	netConn := &MockNetworkWriter{}
	netWriter := ipc.NewWriter(netConn, ipc.WithSchema(schema))
	defer netWriter.Close()

	// ---------------------------------------------------------
	// 2. 文件翻转控制变量 (修改 Writer 类型)
	// ---------------------------------------------------------
	var (
		currentFile *os.File
		// 【修改点 1】类型变为 pqarrow.FileWriter
		currentFileWriter *pqarrow.FileWriter
		err               error
		fileIndex         int
		ticker            = time.NewTicker(1 * time.Second)
	)
	defer ticker.Stop()

	// 辅助函数：创建一个新的 Parquet 文件
	rotateFile := func() error {
		// a. 关闭旧文件 (写入 Parquet Footer)
		if currentFileWriter != nil {
			if err := currentFileWriter.Close(); err != nil {
				return fmt.Errorf("关闭旧 Parquet Writer 失败: %w", err)
			}
			fmt.Printf("--- 文件 data_%d.parquet 已归档 (写入 Footer) ---\n", fileIndex)
		}

		// b. 创建新文件
		fileIndex++
		filename := fmt.Sprintf("data_%d.parquet", fileIndex) // 后缀名改为 .parquet
		currentFile, err = os.Create(filename)
		if err != nil {
			return err
		}

		// c. 创建新的 Parquet Writer
		// 【修改点 2】初始化 Parquet Writer 配置
		// 通常建议开启压缩 (如 Snappy 或 Zstd)
		props := parquet.NewWriterProperties(
			parquet.WithCompression(compress.Codecs.Snappy), // 使用 Snappy 压缩
			parquet.WithCreatedBy("Go Arrow Example"),
		)

		// 使用 pqarrow 库创建 Writer，它能直接接受 Arrow Record
		currentFileWriter, err = pqarrow.NewFileWriter(
			schema,
			currentFile,
			props,
			pqarrow.DefaultWriterProps(),
		)
		if err != nil {
			return err
		}

		fmt.Printf(">>> 创建新文件: %s\n", filename)
		return nil
	}

	// 首次启动先创建一个文件
	if err := rotateFile(); err != nil {
		log.Fatal(err)
	}

	// ---------------------------------------------------------
	// 3. 数据循环
	// ---------------------------------------------------------
	b := array.NewRecordBuilder(pool, schema)
	defer b.Release()

	for i := 0; i < 50; i++ {
		// 检查翻转
		select {
		case <-ticker.C:
			if err := rotateFile(); err != nil {
				log.Fatal(err)
			}
		default:
		}

		// 生成数据
		b.Field(0).(*array.TimestampBuilder).Append(arrow.Timestamp(time.Now().UnixNano()))
		b.Field(1).(*array.Float64Builder).Append(float64(i) * 1.5)
		rec := b.NewRecord()

		// A. 写入网络 (Arrow IPC)
		if err := netWriter.Write(rec); err != nil {
			log.Printf("网络发送错误: %v", err)
		}

		// B. 写入文件 (Parquet)
		// 【修改点 3】调用 Parquet Writer 的 Write 方法
		// 注意：虽然接口一样，但底层逻辑不同。这里每次 Write 可能会生成一个 RowGroup，
		// 或者根据 Writer 缓冲策略处理。对于 Parquet，RecordBatch 最好大一点。
		if err := currentFileWriter.Write(rec); err != nil {
			log.Printf("Parquet 写入错误: %v", err)
		}

		rec.Release()
		time.Sleep(100 * time.Millisecond)
	}

	// 清理
	if currentFileWriter != nil {
		currentFileWriter.Close()
		currentFile.Close()
	}
}
