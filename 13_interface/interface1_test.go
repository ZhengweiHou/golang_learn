package hzwinterface

import (
	"fmt"
	"testing"
)

func TestInt1(t *testing.T) {
	var timp TIntface1
	timp = &TImp{
		Name: "hhh",
	}

	// 判断timp是否实现了TIntface2
	t1, ok := timp.(TIntface2) // 注意此处的timp的定义需要是个接口
	if ok {
		t1.Test2()
	}
	fmt.Printf("%p,%p\n", timp, t1)
}

type TIntface1 interface {
	Test1()
}

type TIntface2 interface {
	Test2()
}

type TImp struct {
	Name string
}

func (t *TImp) Test1() {
	fmt.Printf("%s Test1\n", t.Name)
}

func (t *TImp) Test2() {
	fmt.Printf("%s Test2\n", t.Name)
}
