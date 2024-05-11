package base

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestLimit1(t *testing.T) {
	limit := rate.NewLimiter(1, 5)

	ctx := context.Background()
	for i := 0; i < 10; i++ {
		limit.Wait(ctx)
		fmt.Printf("%d  %v\n", i, time.Now())
	}
	fmt.Println("=========")
	time.Sleep(3 * time.Second)
	for i := 0; i < 10; i++ {
		limit.Wait(ctx)
		fmt.Printf("%d  %v\n", i, time.Now())
	}

	limit.SetBurst(5)
	limit.SetLimit(2)
	fmt.Println("=========")
	time.Sleep(3 * time.Second)
	for i := 0; i < 10; i++ {
		limit.Wait(ctx)
		fmt.Printf("%d  %v\n", i, time.Now())
	}
}
