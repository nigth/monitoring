package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Count = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_count",
		Help: "App Request Count",
	},
		[]string{"app_name", "method", "endpoint", "http_status"},
	)
	// Latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	// 	Name:    "request_latency_seconds",
	// 	Help:    "Request latency",
	// 	Buckets: prometheus.DefBuckets([]float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 7.5, 10}),
	// },
	// 	[]string{"app_name", "endpoint"},
	// )
)

func init() {
	prometheus.MustRegister(Count)
	//prometheus.MustRegister(Latency)
}

// func StartTimer() {
// 	http.Request.StartTime = time.Now()
// }

// func StopTimer() {
// 	RespTime = time.Now() - http.Request.StartTime
// 	Latency.
// }

// func StartTime() {
// 	var Start = time.Now()
// }

// func MeasureTime() {
// 	var Start = time.Now()
// 	Latency.Observer(time.Since(Start).Seconds())
// }

func PostCount() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		Count.With(prometheus.Labels{"app_name": "post-srv", "method": r.Method,
			"endpoint": r.Host, "http_status": r.Response.Status}).Inc()
	})
	return
}

// func PostHist() {
// 	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
// 		Latency.With(prometheus.Labels{"app_name": "post-srv", "endpoint": r.Host})
// 	})
// }

func Output() {
	http.Handle("/metrics", promhttp.Handler())
}
