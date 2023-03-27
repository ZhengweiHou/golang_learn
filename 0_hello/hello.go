package main

import "fmt"

/*
一个main包有且只有一个main函数，main函数是程序的执行入口
*/

func main() { // 函数体的左括号需要函数名在同一行
	fmt.Println("hello hzw!")
}

// go build hello.go
// go run hello.go
