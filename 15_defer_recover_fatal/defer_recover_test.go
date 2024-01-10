package main

import (
	"fmt"
	"log"
	"testing"
)

func TestDeferRecover1(t *testing.T) {
	/*
		panic：词义"恐慌"，
		recover："恢复"
		go语言利用panic()，recover()，实现程序中的极特殊的异常的处理
			panic(),让当前的程序进入恐慌，中断程序的执行
			recover(),让程序恢复，必须在defer函数中执行
	*/
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("man defer recover msg:%s\n", msg)
		}
	}()

	funA()
	defer fmt.Println("defer after call funcA")

	funB()
	defer fmt.Println("defer after call funcB")

	fmt.Println("main end")
}

func funA() {
	fmt.Println("run funcA")
}

func funB() { //外围函数
	fmt.Println("start funB")
	defer fmt.Println("defer funB 1")

	for i := 1; i <= 10; i++ {
		fmt.Printf("funcB for with %d\n", i)
		if i == 5 {
			//让程序中断
			panic("funB panic")
		}
	}
	//当外围函数的代码中发生了运行恐慌，只有其中所有的已经defer的函数全部都执行完毕后，该运行恐慌才会真正被扩展至调用处。
	defer fmt.Println("defer funB 2") // 此处被panic中断，不会执行
}

func funC() {
	log.Fatal("funC fatal") // fatal会直接结束程序，不会被recover捕获
}
