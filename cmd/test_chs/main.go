package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sonirico/go-hyperliquid"

	dr "indexer/data_retriver" // 根据模块名称导入
)

var fill = hyperliquid.Fill{}

//go:embed users.txt
var usersTxt string

func main() {
	go pprofGoroutine()
	// 解析命令行参数
	var workers int
	var delayMs int
	flag.IntVar(&workers, "workers", 5, "number of worker goroutines")
	flag.IntVar(&delayMs, "delay", 0, "delay in milliseconds between requests")
	flag.Parse()

	// 初始化客户端

	// 从嵌入的 users.txt 获取地址列表
	users := strings.Split(strings.TrimSpace(usersTxt), "\n")

	// 过滤掉空地址
	var validUsers []string
	for _, user := range users {
		if user != "" {
			validUsers = append(validUsers, user)
		}
	}

	if len(validUsers) == 0 {
		log.Fatal().Msg("No valid users found in users.txt")
	}

	var count atomic.Int32

	// 启动一个goroutine定期打印QPS
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			<-ticker.C
			qps := float64(count.Load()) / 5.0
			fmt.Printf("QPS: %.2f\n", qps)
			count.Store(0)
		}
	}()

	// 创建等待组以管理多个goroutine
	var wg sync.WaitGroup

	// 启动指定数量的工作goroutine
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			// 每个goroutine从不同的位置开始处理地址
			for j := 0; ; j++ {
				address := validUsers[(workerID+j)%len(validUsers)]

				// 对每个地址发起 UserState 请求
				_, err := dr.FetchUserState(address)
				if err != nil {
					log.Error().Err(err).Str("address", address).Msg("Failed to get user state")
				} else {
					count.Add(1)
				}

				// 如果设置了延迟，则等待
				if delayMs > 0 {
					time.Sleep(time.Duration(delayMs) * time.Millisecond)
				}
			}
		}(i)
	}

	// 等待所有goroutine完成（实际上不会完成，因为是无限循环）
	wg.Wait()
}

func pprofGoroutine() {
	profileFileName := "cpu_profile_after_5s.prof"
	f, err := os.Create(profileFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create profile file: %v\n", err)
		os.Exit(1)
	}
	// 确保文件最终会被关闭
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Failed to close profile file: %v\n", closeErr)
		} else {
			fmt.Printf("Profile data saved to %s\n", profileFileName)
		}
	}()

	fmt.Println("Waiting for 5 seconds before starting profiling...")
	time.Sleep(5 * time.Second)

	// 4. 开始 CPU profiling，数据写入到文件 f
	fmt.Println("Starting CPU profiling for 10 seconds...")
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Fprintf(os.Stderr, "Could not start CPU profile: %v\n", err)
		os.Exit(1)
	}

	// 5. 录制 5 秒钟的 profile 数据
	time.Sleep(10 * time.Second)

	// 6. 停止 profiling
	pprof.StopCPUProfile()
	fmt.Println("Finished CPU profiling.")
	os.Exit(0)
}
