package base

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGroup1(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(3 * time.Second)
		}()
	}
	fmt.Println(time.Now())
	wg.Wait()
	fmt.Println(time.Now())
}
