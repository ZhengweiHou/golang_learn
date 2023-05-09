package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("hello start..")
	quit := make(chan os.Signal)
	// signal.Notify(quit, os.Kill)
	signal.Notify(quit) // 不指定信号，则监听所有信号
	s := <-quit
	fmt.Println("Got signal:", s)
	time.Sleep(time.Second)
	fmt.Println("hello End")
}
