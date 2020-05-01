package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	CounterVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_count",
		Help: "App Request Count",
	}, []string{"app_name", "method", "endpoint", "http_status"},
	)
	HistogramVec = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "request_latency_seconds",
		Help: "Request latency",
	}, []string{"app_name", "endpoint"},
	)
)

func RegisterMetrics() {
	prometheus.Register(CounterVec)
	prometheus.Register(HistogramVec)
}

func Init(req *http.Request) {
	start := time.Now()
	status := ""
	endpoint := req.URL.Path
	serName := "post-srv"
	method := req.Method

}

func Collect() {
	Init(req)
	CounterVec.WithLabelValues(serName, method, endpoint, status).Inc()
	HistogramVec.WithLabelValues(serName, endpoint).Observe(time.Since(start).Seconds())
}

func StatusCollect() {
	Init(req)
	status = req.resp.Status
}

func Output() {
	http.Handle("/metrics", promhttp.Handler())
}
