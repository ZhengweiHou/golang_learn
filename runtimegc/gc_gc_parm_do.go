package runtimegc

import (
	"fmt"
	"net/http"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"sync"
	"time"
)

func DoGcTest() {
	//创建trace文件
	tf, _ := os.Create("/tmp/trace.out")
	defer tf.Close()
	trace.Start(tf)
	defer trace.Stop()

	cf, _ := os.Create("/tmp/cpuprofile.out")
	defer cf.Close()
	pprof.StartCPUProfile(cf)
	defer pprof.StopCPUProfile()

	startt := time.Now()

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	fmt.Println("hello world")
	wg := new(sync.WaitGroup)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go leakyFunction(wg, i)
	}
	wg.Wait()
	fmt.Printf("%v\n", time.Now().Sub(startt))
}

func leakyFunction(wg *sync.WaitGroup, gindex int) {
	defer wg.Done()
	for n := 0; n < 100; n++ {
		nt := time.Now()
		s := make([]string, 3)
		for i := 0; i < 1000000; i++ {
			s = append(s, "magical pandas")
			// if (i % 100000) == 0 {
			// time.Sleep(1 * time.Millisecond)
			// }
		}
		fmt.Printf("%d-%d %v\n", gindex, n, time.Now().Sub(nt))
		// time.Sleep(time.Second)
	}
}
