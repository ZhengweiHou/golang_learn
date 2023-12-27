package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"sync"
)

func main() {

	// go tool trace trace.out 启动trace分析

	// runtime.GOMAXPROCS(runtime.NumCPU()) // 设置P队列的最大数量 (理解GMP模型)
	runtime.GOMAXPROCS(4)
	debug.SetMaxThreads(5)

	//创建trace文件
	f, err := os.Create("/tmp/trace.out")
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

	//main
	fmt.Println("Hello World")

	var wg sync.WaitGroup

	numTasks := 120 // 启动n个 goroutine

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go performComputation(&wg)
	}

	wg.Wait()

}

func performComputation(wg *sync.WaitGroup) {
	defer wg.Done()

	const iterations = 10000
	var result float64

	for i := 0; i < iterations; i++ {
		result += float64(i)
	}

	fmt.Println(result)
}
