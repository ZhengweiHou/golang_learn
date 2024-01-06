package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChannel1(t *testing.T) {
	c1 := make(chan string, 3)

	c1 <- "hhh"
	c1 <- "zzz"
	c1 <- "www"

	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)

	c1 <- "111"

	fmt.Println(len(c1)) // 1
	_, ok := <-c1
	fmt.Println(ok)      // true
	fmt.Println(len(c1)) // 0

	c1 <- "222"

	close(c1) // chanl中有数据，close关不掉

	fmt.Println(len(c1)) // 1
	_, ok = <-c1
	fmt.Println(ok)      // false
	fmt.Println(len(c1)) // 0

}

func TestChan2(t *testing.T) {
	ch := make(chan bool, 3)
	wg := &sync.WaitGroup{}

	startT := time.Now()

	fmt.Printf("start:%v\n", time.Now())
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(ch chan bool) {
			defer wg.Done()
			ch <- true
			time.Sleep(time.Millisecond * 200)
			<-ch
			fmt.Println(time.Now())
		}(ch)
	}
	wg.Wait()
	time := time.Now().Sub(startT)
	fmt.Println("take time:", time)
}
