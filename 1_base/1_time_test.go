package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_time1(t *testing.T) {
	time1 := time.Now()

	time.Sleep(time.Second * 2)

	fmt.Println(time.Second)
	fmt.Println(time1.Sub(time.Now()))
	fmt.Println(time.Now().Sub(time1))
	fmt.Println("毫秒数：", time1.UnixMilli())

	if time1.Sub(time.Now()) < time.Second {
		fmt.Println("1111")
	}
	if time.Now().Sub(time1) > time.Second {
		fmt.Println("2222")
	}
}
