package hzwinterface

import (
	"fmt"
	"testing"
)

func TestIntPointRef(t *testing.T) {
	taskMap := make(map[string]TIP1)
	taskMap["temp"] = &htip1{}

	h1, _ := taskMap["temp"]
	h2 := h1
	h2.SetName("hzw")
	h2.Name()
	h1.Name()
}

type TIP1 interface {
	SetName(name string)
	Name()
}

type htip1 struct {
	name string
}

func (h *htip1) Name() {
	fmt.Println(h.name)
}
func (h *htip1) SetName(name string) {
	h.name = name
}
