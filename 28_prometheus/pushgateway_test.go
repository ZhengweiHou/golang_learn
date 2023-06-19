package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
)

func TestPushgateway(t *testing.T) {
	cnt := prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "hhh", Help: "hhh test"},
		[]string{"hcode"},
	)

	prometheus.MustRegister(cnt)

	// pushgateway 接入方式
	pusher := push.New("localhost:9091", "hzwapp")
	pusher.Collector(cnt) // 关联指标收集器
	go func() {
		for i := 0; i < 55; i++ {
			err := pusher.Push()
			if err != nil {
				fmt.Println(err)
				// fmt.Errorf("err:%v", err)
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		codes := []string{"h", "z", "w"}
		for i := 0; i < 50; i++ {
			time.Sleep(time.Second)
			cnt.WithLabelValues(codes[rand.Intn(3)]).Inc()
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)

}

func Test2(t *testing.T) {
	pusher := push.New("localhost:9091", "hzwapp")
	cnt := prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "hhh", Help: "hhh test"},
		[]string{"hcode"},
	)
	cnt.WithLabelValues("hhh").Inc()

	pusher.Collector(cnt)

	// registry := prometheus.NewRegistry()  // 向创建一个自定义的register
	// registry.MustRegister(httpRequestCounter, concurrentHttpRequestsGauge, concurrentHttpRequestsGauge, summary)  // 向register中注册多个meterics

	err := pusher.Push()
	// pusher.Add()

	fmt.Println(err)
}
