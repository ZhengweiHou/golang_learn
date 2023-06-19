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

// create table 
