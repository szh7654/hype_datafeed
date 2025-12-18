package metric

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

var requestCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Total number of HTTP requests.",
})