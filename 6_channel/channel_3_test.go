package main

import (
	"fmt"
	"testing"
	"time"
)

// 
func TestChannel3(t *testing.T) {
	c1 := make(chan []interface{}, 1)

	go func() {
		record := <-c1
		time.Sleep(time.Second)
		fmt.Println("=go=", record)
	}()

	var record = []interface{}{}
	record = append(record, 1)
	record = append(record, 2)
	c1 <- record
	fmt.Println("=1", record)
	// record = []interface{}{}
	record = append(record, 11)
	fmt.Println("=2", record)
	time.Sleep(time.Second * 2)

}

func TestChannel3_2(t *testing.T) {
	c1 := make(chan []interface{}, 1)

	go func() {
		record := <-c1
		time.Sleep(time.Second)
		fmt.Println("=go=", record)
	}()

	var record = []interface{}{}
	record = append(record, 1)
	record = append(record, 2)
	// record[0] = 1
	// record[1] = 2
	c1 <- record
	fmt.Println("=1", record)
	// record = make([]interface{}, 2) // 重新创建新的切片不会影响协程中读取的数据
	// record = []interface{}{}
	record = append(record, 11)
	// record[0] = 11 // 会影响到协程中读取的数据
	fmt.Println("=2", record)
	time.Sleep(time.Second * 2)

	// =1 [1 2]
	// =2 [<nil> <nil>]
	// == [1 2]
}
