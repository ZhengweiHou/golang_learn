package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	// _ "github.com/uber/go-torch"
)

func Counter(wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 100000000000; i++ {
		// time.Sleep(time.Millisecond * 200)
		counter++
	}
	wg.Done()
}

func main() {
	flag.Parse()

	//远程获取pprof数据
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8588", nil)
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Counter(&wg)
	}
	wg.Wait()

	// sleep 10mins, 在程序退出之前可以查看性能参数.
	// time.Sleep(60 * time.Second)
}
