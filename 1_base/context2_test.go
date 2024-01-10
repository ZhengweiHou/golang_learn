package base

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext_1(t *testing.T) {
	mct, mcancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			n := 1
			for {
				select {
				case <-time.After(time.Second):
					fmt.Printf("Worker%d run:%d\n", i, n)
					n++
				case <-ctx.Done():
					fmt.Printf("Worker%d down\n", i)
					return
				}
			}
		}(i, mct)
	}

	mct.Done()

	time.Sleep(5 * time.Second)
	mcancel()
	time.Sleep(2 * time.Second)

}

func TestContext_2(t *testing.T) {
	mct := context.Background()

	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			n := 1
			for {
				select {
				case <-time.After(time.Second):
					fmt.Printf("Worker%d run:%d\n", i, n)
					n++
				case <-ctx.Done():
					fmt.Printf("Worker%d down\n", i)
					return
				}
			}
		}(i, mct)
	}

	time.Sleep(5 * time.Second)
	time.Sleep(2 * time.Second)
}
