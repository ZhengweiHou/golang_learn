package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestFatalInGoroutine(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)
	go funFatal1(&wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Printf("normal func for with %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()

}

// 产生Fatal
func funFatal1(wg *sync.WaitGroup) {

	defer func() {
		fmt.Println("funcFatal defer")
		wg.Done()
	}()
	for i := 1; i <= 10; i++ {
		fmt.Printf("funcFatal for with %d\n", i)
		if i == 5 {
			log.Fatal("funFatal fatal") // Fatal 会直接中断程序
		}
		time.Sleep(time.Second)
	}
}
