package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	dr "indexer/data_retriver" // 根据模块名称导入
)

//go:embed users.txt
var usersTxt string

func main() {
	// 解析命令行参数
	var workers int
	var delayMs int
	flag.IntVar(&workers, "workers", 1, "number of worker goroutines")
	flag.IntVar(&delayMs, "delay", 0, "delay in milliseconds between requests")
	flag.Parse()

	// 初始化客户端
	client := dr.NewClient(dr.LocalAPIURL, false)

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

	var count int64
	var mutex sync.Mutex

	// 启动一个goroutine定期打印QPS
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			<-ticker.C
			mutex.Lock()
			qps := float64(count) / 5.0
			fmt.Printf("QPS: %.2f\n", qps)
			count = 0 // 重置计数器
			mutex.Unlock()
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
				_, err := client.UserState(address)
				if err != nil {
					log.Error().Err(err).Str("address", address).Msg("Failed to get user state")
				} else {
					mutex.Lock()
					count++
					mutex.Unlock()
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
