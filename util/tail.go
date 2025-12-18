package util

import (
	"context"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/nxadm/tail"
	"github.com/rs/zerolog/log"
)

// HourlyTailer 结构体
type HourlyTailer struct {
	baseDir string
	Lines   chan string
}

func NewTailer(baseDir string, ctx context.Context, wg *sync.WaitGroup) *HourlyTailer {
	h := &HourlyTailer{
		baseDir: baseDir,
		// 缓冲区稍微大一点，防止双倍流量写入时阻塞
		Lines: make(chan string, 2048),
	}
	wg.Add(1)
	go h.managerLoop(ctx, wg)
	return h
}

// managerLoop 负责监控时间，调度具体的 tail 任务
// 它只负责“什么时候开启新文件的 tail”，不负责读取数据
func (h *HourlyTailer) managerLoop(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	// 定时检查新文件
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	yymmdd := ""
	hour := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			t := time.Now()
			currentYymmdd := t.Format("20060102")
			currentHour := t.Hour()
			if currentYymmdd != yymmdd || currentHour != hour {
				h.startTailWorker(currentYymmdd, currentHour, ctx, wg)
				yymmdd = currentYymmdd
				hour = currentHour
			}
		}
	}
}

func (h *HourlyTailer) startTailWorker(yymmdd string, hour int, ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		filePath := filepath.Join(h.baseDir, yymmdd, strconv.Itoa(hour))

		t, err := tail.TailFile(filePath, tail.Config{
			ReOpen:    false,
			Follow:    true,
			MustExist: false,
			Poll:      false,
			Logger:    tail.DiscardingLogger,
		})

		if err != nil {
			log.Err(err).Msgf("start tail %s error", filePath)
			return
		}
		defer t.Stop()
		log.Info().Msgf("Read Worker started: %s", filePath)

		idleTime := time.NewTimer(time.Minute * 1)

		for {
			select {
			case <-ctx.Done():
				return
			case line, ok := <-t.Lines:
				idleTime.Reset(time.Minute * 1)
				if !ok {
					log.Info().Msgf("Tail File %s closed", filePath)
					return
				}
				if line.Err != nil {
					log.Err(line.Err).Msgf("Read line error: %s", line.Text)
					continue
				}
				h.Lines <- line.Text
			case <-idleTime.C:
				log.Info().Msgf("Worker %s is idle. check is hour is past", filePath)
				if hour != time.Now().Hour() {
					return
				}
			}
		}
	}()
}

func (h *HourlyTailer) getLogPath(t time.Time) string {
	dateDir := t.Format("20060102")
	hourFile := strconv.Itoa(t.Hour())
	return filepath.Join(h.baseDir, dateDir, hourFile)
}
