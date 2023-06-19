package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// func TestCounter(t *testing.T) {

// }

func TestGague(t *testing.T) {
	// CommonGauge
	var CommonGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s%s", namePrefix, "common"),
		},
	)

	// FuncGauge
	var funcGaugeTotalCount int64
	var FuncGauge = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s%s", namePrefix, "func"),
	}, func() float64 {
		deltaVal := rand.Int63n(5) // 模拟变动步长
		idx := rand.Intn(3)        // 模拟增加还是减少 0减少，1和2增加，整体保持一个增长趋势图会好看点
		var newNum int64
		if idx == 0 {
			newNum = atomic.AddInt64(&funcGaugeTotalCount, -deltaVal)
		} else {
			newNum = atomic.AddInt64(&funcGaugeTotalCount, deltaVal)
		}
		return float64(newNum)
	})

	// VecGauge
	var VecGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s%s", namePrefix, "vec"),
		}, []string{"status"})

	prometheus.MustRegister(CommonGauge, FuncGauge, VecGauge)

	go func() {
		codes := []string{"h", "z", "w"}
		for i := 0; i < 50; i++ {
			time.Sleep(time.Second)
			VecGauge.WithLabelValues(codes[rand.Intn(3)]).Inc()
			CommonGauge.Add(float64(i))
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
	time.Sleep(time.Second)
}

// 并发测试
func TestConcurrency(t *testing.T) {
	c := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hzwtest",
	})
	// c.Inc()

	prometheus.MustRegister(c)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	for i := 0; i < 10; i++ {
		go func() {
			wg.Wait()
			for n := 0; n < 10; n++ {
				c.Inc()
			}
		}()
	}
	wg.Done()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
	time.Sleep(time.Second)
}

const (
	namePrefix = "hzw_the_gague_of_"
	subSys     = "hzw_client"
	nameSpace  = "hzwgolanglearn"
)
