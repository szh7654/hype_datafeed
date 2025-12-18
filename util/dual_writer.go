package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/apache/arrow/go/v18/arrow"
	"github.com/apache/arrow/go/v18/arrow/array"
	"github.com/apache/arrow/go/v18/arrow/ipc"
	"github.com/apache/arrow/go/v18/arrow/memory"
	"github.com/apache/arrow/go/v18/parquet"
	"github.com/apache/arrow/go/v18/parquet/compress"
	"github.com/apache/arrow/go/v18/parquet/pqarrow"
)

// Config 配置项
type Config struct {
	Schema           *arrow.Schema // Arrow Schema 定义
	OutputDir        string        // 文件输出目录
	FilePrefix       string        // 文件名前缀 (如 "sensor_data")
	RotationInterval time.Duration // 轮转时间间隔
	NetworkWriter    io.Writer     // 网络输出流 (可选)
}

// Writer 封装后的写入器
type Writer struct {
	config Config
	pool   memory.Allocator

	// 内部通信管道
	recordCh chan arrow.Record
	doneCh   chan struct{}

	// 记录构建器 (用于将 struct 转为 arrow)
	builder *array.RecordBuilder
}

// New 创建一个新的 DualWriter
func NewDualWriter(cfg Config) (*Writer, error) {
	if cfg.Schema == nil {
		return nil, fmt.Errorf("schema cannot be nil")
	}
	if err := os.MkdirAll(cfg.OutputDir, 0755); err != nil {
		return nil, fmt.Errorf("create dir failed: %w", err)
	}

	pool := memory.NewGoAllocator()
	w := &Writer{
		config:   cfg,
		pool:     pool,
		recordCh: make(chan arrow.Record, 100), // 缓冲 100 个批次
		doneCh:   make(chan struct{}),
		builder:  array.NewRecordBuilder(pool, cfg.Schema),
	}

	// 启动后台处理协程
	go w.run()

	return w, nil
}

// Write 接收 []struct，使用反射转换为 Arrow Record 并异步写入
// data 必须是 slice 类型，且字段顺序需与 Schema 一致
func (w *Writer) Write(data interface{}) error {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		return fmt.Errorf("data must be a slice")
	}

	length := val.Len()
	if length == 0 {
		return nil
	}

	// 1. 使用反射将 struct 数据填充到 RecordBuilder
	// 注意：为了性能，这里假设 struct 字段顺序与 Schema 字段顺序一致
	// 且类型匹配 (Int64 -> Int64, Float64 -> Float64)
	fields := w.config.Schema.Fields()
	for i := 0; i < length; i++ {
		item := val.Index(i)

		for j, field := range fields {
			// 获取 struct 的第 j 个字段
			fieldVal := item.Field(j)

			// 根据 Schema 类型追加数据 (这里仅演示常用类型，可扩展)
			switch field.Type.ID() {
			case arrow.INT64:
				w.builder.Field(j).(*array.Int64Builder).Append(fieldVal.Int())
			case arrow.FLOAT64:
				w.builder.Field(j).(*array.Float64Builder).Append(fieldVal.Float())
			case arrow.STRING:
				w.builder.Field(j).(*array.StringBuilder).Append(fieldVal.String())
			case arrow.BOOL:
				w.builder.Field(j).(*array.BooleanBuilder).Append(fieldVal.Bool())
			// TODO: 添加 Timestamp 等其他类型的支持
			default:
				// 复杂类型建议通过自定义接口处理，避免反射过于复杂
			}
		}
	}

	// 2. 生成 Record
	rec := w.builder.NewRecord() // 这是一个轻量级操作，数据是指针引用或 copy
	
	// 3. 发送到后台处理 (非阻塞，除非管道满)
	select {
	case w.recordCh <- rec:
		return nil
	case <-w.doneCh:
		rec.Release()
		return fmt.Errorf("writer is closed")
	}
}

// Close 关闭写入器
func (w *Writer) Close() {
	close(w.recordCh)
	<-w.doneCh // 等待后台任务完全退出
	w.builder.Release()
}

// run 后台主循环：处理 IO 和文件轮转
func (w *Writer) run() {
	defer close(w.doneCh)

	// 初始化 IPC Stream Writer (网络)
	var streamWriter *ipc.Writer
	if w.config.NetworkWriter != nil {
		streamWriter = ipc.NewWriter(w.config.NetworkWriter, ipc.WithSchema(w.config.Schema))
		defer streamWriter.Close()
	}

	// 初始化 Parquet Writer 状态
	var (
		fileIndex    int
		currFile     *os.File
		currPqWriter *pqarrow.FileWriter
		ticker       = time.NewTicker(w.config.RotationInterval)
	)
	defer ticker.Stop()

	// 内部闭包：执行文件轮转
	rotate := func() error {
		// 关闭旧文件
		if currPqWriter != nil {
			currPqWriter.Close()
			currFile.Close()
			fmt.Printf("--- Log Rotated: data_%d.parquet ---\n", fileIndex)
		}

		// 创建新文件
		fileIndex++
		filename := fmt.Sprintf("%s_%d.parquet", w.config.FilePrefix, fileIndex)
		path := filepath.Join(w.config.OutputDir, filename)

		f, err := os.Create(path)
		if err != nil {
			return err
		}
		currFile = f

		// 创建带压缩的 Parquet Writer
		props := parquet.NewWriterProperties(
			parquet.WithCompression(compress.Codecs.Snappy),
		)
		pqW, err := pqarrow.NewFileWriter(w.config.Schema, currFile, props, pqarrow.DefaultWriterProps())
		if err != nil {
			return err
		}
		currPqWriter = pqW
		return nil
	}

	// 启动时先创建第一个文件
	if err := rotate(); err != nil {
		fmt.Printf("Initial rotation failed: %v\n", err)
		return
	}

	// 事件循环
	for {
		select {
		// A. 接收到数据
		case rec, ok := <-w.recordCh:
			if !ok {
				// Channel 关闭，退出前清理当前文件
				if currPqWriter != nil {
					currPqWriter.Close()
					currFile.Close()
				}
				return
			}

			// 1. 写网络 (IPC Stream)
			if streamWriter != nil {
				if err := streamWriter.Write(rec); err != nil {
					fmt.Printf("Net write error: %v\n", err)
				}
			}

			// 2. 写文件 (Parquet)
			if currPqWriter != nil {
				if err := currPqWriter.Write(rec); err != nil {
					fmt.Printf("File write error: %v\n", err)
				}
			}

			// 重要：释放 Record 内存 (因为是通过 Channel 传递过来的)
			rec.Release()

		// B. 时间到，轮转文件
		case <-ticker.C:
			if err := rotate(); err != nil {
				fmt.Printf("Rotate error: %v\n", err)
			}
		}
	}
}
