package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_string1(t *testing.T) {
	var str1 string

	fmt.Printf("len:%v\n", len(str1))

	str1 = "123123123"

	fmt.Println(strings.LastIndex(str1, "2"))
	fmt.Println(strings.Index(str1, "2"))
	fmt.Println(strings.Count(str1, "2"))

	var str2 = "hello"

	if 1 == 1 {
		str2 = "hehe"
	}

	fmt.Println(str2)

}

func Test_string2(t *testing.T) {
	input := "hello:123"
	parts := strings.Split(input, ":")
	fmt.Println(len(parts))

	input = "hello"
	parts = strings.Split(input, ":")
	fmt.Println(len(parts))

	input = ""
	parts = strings.Split(input, ":")
	fmt.Println(len(parts))

}

func Test_string3(t *testing.T) {
	str1 := "hello"
	str2 := "Hello"
	fmt.Println(strings.Compare(str1, str2))
}
