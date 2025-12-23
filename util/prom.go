package util

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	// Register Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	// Start HTTP server in a separate goroutine
	// pprof endpoints are automatically registered by importing _ "net/http/pprof"
	go http.ListenAndServe(":2112", nil)
}
