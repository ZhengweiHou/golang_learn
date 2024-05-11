package hzwinterface

import (
	"fmt"
	"testing"
)

func TestInt2(t *testing.T) {

	var h IH
	h = &H1{
		BH{
			name: "hzw",
		},
	}
	h.Hello()
}

// 接口
type IH interface {
	Hello()
	doHello()
}

// BH 持有IH
type BH struct {
	IH
	name string
}

// BH 实现了Hello()
func (h *BH) Hello() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("BH Hello() %T err:%v", err, err)
		}
	}()
	fmt.Println(h.name)
	h.doHello() // 试图在父类调用子类的方法，但此处调用的是BH中IH的doHello，此处会空指针
}

// H1 是 BH 的子类，实现doHello
type H1 struct {
	BH
}

func (h *H1) doHello() {
	fmt.Println("我是H1")
}
