package runtimetest

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// 获取当前GOMAXPROCS 参数<=0时，返回当前的值，>0时，会设置GOMAXPROCS
	gps := runtime.GOMAXPROCS(0)
	ncpu := runtime.NumCPU()
	fmt.Printf("gps:%d,ncpu:%d\n", gps, ncpu)

	go func() { time.Sleep(time.Second) }()
	go func() { time.Sleep(time.Second) }()
	go func() { time.Sleep(time.Second) }()

	ngo := runtime.NumGoroutine()
	// 设置M的数量
	maxT := debug.SetMaxThreads(7)
	fmt.Printf("ngo:%d,maxT:%d\n", ngo, maxT)

}
