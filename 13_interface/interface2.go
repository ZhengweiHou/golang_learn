package main

import (
	"fmt"
)

// func TestXxx1(t *testing.T) {
func main() {

	var h IH
	h = &H1{
		BaseH{
			name: "hzw",
		},
	}

	h.Hello()

}

type IH interface {
	Hello()
	doHello()
}
type BaseH struct {
	IH
	name string
}
type BaseH2 struct {
	Ih   IH
	name string
}

func (h *BaseH) Hello() {
	fmt.Println(h.name)
	h.IH.doHello()
}

type H1 struct {
	BaseH
}

func (h *H1) doHello() {
	fmt.Println("我是H1")
}
