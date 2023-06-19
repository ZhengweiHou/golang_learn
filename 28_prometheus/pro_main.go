package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// 定义自定义的指标
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests.",
		},
	)
)

func init() {
	// 注册自定义的指标
	prometheus.MustRegister(requestsTotal)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 自定义指标计数
		requestsTotal.Inc()
		w.Write([]byte("Hello, world!"))
	})
	http.ListenAndServe(":8080", nil)
}
