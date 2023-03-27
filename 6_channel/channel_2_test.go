package main

import (
	"log"
	"sync"
	"testing"
	"time"
)

func Test1(t *testing.T) {

	var throttleChannel chan bool
	throttleChannel = make(chan bool, 2)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			throttleChannel <- true // 占住一个位置
			log.Println("==")
			time.Sleep(time.Duration(2) * time.Second)
			<-throttleChannel // 释放位置
		}()
	}

	wg.Wait()

}
