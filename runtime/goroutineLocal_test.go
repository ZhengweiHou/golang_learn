package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGroutLocal(t *testing.T) {
	// 创建一个sync.Pool
	pool := &sync.Pool{
		New: func() interface{} {
			// 在需要时，为每个goroutine创建新的变量
			return "Default Value"
		},
	}

	// 启动多个goroutine
	for i := 0; i < 3; i++ {
		go func(id int) {
			// 从池中获取goroutine-local变量
			value := pool.Get()
			value = 123
			defer pool.Put(value)

			for n := 0; n < 5; n++ {
				value := pool.Get()
				time.Sleep(time.Second)
				// 在goroutine内使用变量
				fmt.Printf("Goroutine %d: %s\n", id, value)
			}
		}(i)
	}

	// 等待所有goroutine完成
	fmt.Scanln()
}
