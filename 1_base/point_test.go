package main

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var a *string
	fmt.Print(a)
	if a != nil {
		fmt.Print(a)
	}
}

var h *hBean

func TestPoint2(t *testing.T) {
	h = &hBean{
		name: "hzw",
	}

	htemp := *h             // 从指针获取实际值的拷贝
	htemp.name = "123"      // 拷贝的修改不会影响原值
	fmt.Println(h.name)     // 输出:hzw
	fmt.Println(htemp.name) // 输出:123

}

type hBean struct {
	name string
}
