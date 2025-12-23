package util

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// LogCounterHook 实现zerolog.Hook接口来捕获日志事件
type LogCounterHook struct {
	debugCounter prometheus.Counter
	infoCounter  prometheus.Counter
	warnCounter  prometheus.Counter
	errorCounter prometheus.Counter
	fatalCounter prometheus.Counter
	panicCounter prometheus.Counter
}

// Run 实现Hook接口
func (h *LogCounterHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	switch level {
	case zerolog.DebugLevel:
		h.debugCounter.Inc()
	case zerolog.InfoLevel:
		h.infoCounter.Inc()
	case zerolog.WarnLevel:
		h.warnCounter.Inc()
	case zerolog.ErrorLevel:
		h.errorCounter.Inc()
	case zerolog.FatalLevel:
		h.fatalCounter.Inc()
	case zerolog.PanicLevel:
		h.panicCounter.Inc()
	}
}

// 定义日志级别的Prometheus计数器
var (
	logCounters = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "zerolog_log_lines_total",
		},
		[]string{"level"},
	)
)

func init() {
	// 设置全局日志格式，包含代码行号
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Logger.With().Caller().Logger()
	
	// 创建并安装日志计数Hook
	hook := &LogCounterHook{
		debugCounter: logCounters.WithLabelValues("debug"),
		infoCounter:  logCounters.WithLabelValues("info"),
		warnCounter:  logCounters.WithLabelValues("warn"),
		errorCounter: logCounters.WithLabelValues("error"),
		fatalCounter: logCounters.WithLabelValues("fatal"),
		panicCounter: logCounters.WithLabelValues("panic"),
	}
	log.Logger = log.Logger.Hook(hook)
}