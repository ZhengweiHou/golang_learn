package runtimetest

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"testing"
)

func TestTrace(t *testing.T) {

	// go tool trace trace.out 启动trace分析
	// go tool trace /tmp/trace.out
	// runtime.GOMAXPROCS(runtime.NumCPU()) // 设置P队列的最大数量 (理解GMP模型)
	// runtime.GOMAXPROCS(4)
	// debug.SetMaxThreads(5)
	fmt.Println("real GOMAXPROCS:", runtime.GOMAXPROCS(-1))

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

	numTasks := 100 // 启动n个 goroutine

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go performComputation(&wg)
	}

	wg.Wait()
	fmt.Println("real GOMAXPROCS:", runtime.GOMAXPROCS(-1))

}

func performComputation(wg *sync.WaitGroup) {
	defer wg.Done()

	const iterations = 10000000
	var result float64
	var result2 int

	for i := 0; i < iterations; i++ {
		byte1 := make([]byte, 10)
		result += float64(i)
		result2 = len(byte1)
	}

	fmt.Println(result)
	fmt.Println(result2)
}
