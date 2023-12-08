package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// 主线程A，开启若干个子协程B1、B2、B3，当某个子协程发生异常时，需实现主线程和其他所有子协程结束

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 模拟一些工作
	select {
	case <-time.After(time.Duration(id) * time.Second):
		// 模拟一个可能发生异常的情况
		if id == 2 {
			fmt.Printf("Worker %d 发生异常\n", id)
			// 发生异常时取消所有协程
			cancelFunc, ok := ctx.Value("cancelFunc").(context.CancelFunc)
			if ok {
				cancelFunc()
			}
			return
		}
		fmt.Printf("Worker %d 完成工作\n", id)
	case <-ctx.Done():
		// 如果context被取消，表示主线程要求退出
		fmt.Printf("Worker %d 被取消\n", id)
	}
}

func TestContext1(t *testing.T) {
	// 使用WithCancel创建一个可取消的主context
	mainCtx, mainCancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// 启动子协程
	for i := 1; i <= 5; i++ {
		// 为每个子协程创建一个新的context，并将其cancel函数存储到主context中
		ctx, cancel := context.WithCancel(mainCtx)
		mainCtx = context.WithValue(mainCtx, "cancelFunc", cancel)

		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// 主线程等待所有子协程完成或者出现异常
	go func() {
		wg.Wait()
		mainCancel() // 所有子协程完成后取消主context
	}()

	// 主线程等待主context被取消
	<-mainCtx.Done()

	fmt.Println("主线程结束")
}
