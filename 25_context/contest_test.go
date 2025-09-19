package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

const keyID = "id"

func Testmain1(t *testing.T) {
	// 定义上下文值
	rand.Seed(time.Now().Unix())
	ctx := context.WithValue(context.Background(), keyID, rand.Int())
	operation1(ctx)
}

func operation1(ctx context.Context) {
	// do some work

	// we can get the value from the context by passing in the key
	log.Println("operation1 for id:", ctx.Value(keyID), " completed")
	operation2(ctx)
}

func operation2(ctx context.Context) {
	// do some work

	// this way, the same ID is passed from one function call to the next
	log.Println("operation2 for id:", ctx.Value(keyID), " completed")
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("context canceled: %v\n", ctx.Err())
	}

	if ctx.Err() != nil {
		fmt.Printf("context canceled: %v\n", ctx.Err())
	}

	if ctx.Err() == context.Canceled {
		fmt.Println("context canceled")
	}
}
