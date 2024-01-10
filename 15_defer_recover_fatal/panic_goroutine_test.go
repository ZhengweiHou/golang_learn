package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestPanicInGoroutine(t *testing.T) {

	defer func() {
		if msg := recover(); msg != nil { // recover无法捕获子协程中的panic
			fmt.Printf("man defer recover msg:%s\n", msg)
		}
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go funcPanic1(&wg)

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

// 产生Panic
func funcPanic1(wg *sync.WaitGroup) {

	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("funcPanic defer recover msg:%s\n", msg)
		}
	}()

	defer func() {
		fmt.Println("funcPanic defer normal")
		wg.Done()
	}()

	for i := 1; i <= 10; i++ {
		fmt.Printf("funcPainc for with %d\n", i)
		if i == 5 {
			log.Panic("funPanic panic")
		}
		time.Sleep(time.Second)
	}
}
