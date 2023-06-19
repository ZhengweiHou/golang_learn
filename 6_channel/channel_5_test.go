package main

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestChannel5(t *testing.T) {
	chan1 := make(chan bool)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		log.Println("协程开始")
		for i := 0; i < 10; i++ {
			temp, ok := <-chan1
			log.Printf("chan中获取:%v,chan状态:%v", temp, ok)
		}
		wg.Done()
	}()

	chan1 <- true
	chan1 <- true
	time.Sleep(time.Second)
	close(chan1)

	chan1 <- true
	wg.Wait()
}

