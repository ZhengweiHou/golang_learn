package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string
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
