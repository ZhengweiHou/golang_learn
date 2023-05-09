package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChannel4(t *testing.T) {
	var wg sync.WaitGroup

	c1 := make(chan bool, 1)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		c1 <- true
		go func(i int) {
			defer wg.Done()
			fmt.Println("start hello:", i)
			time.Sleep(1 * time.Second)
			fmt.Println("end hello:", i)
			<-c1
		}(i)
	}
	wg.Wait()

	fmt.Println("all END")

}
