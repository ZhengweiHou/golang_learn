package lib1

import (
	"fmt"
	"golang_learn/4_init_and_import/lib2"
)

func Lib1Test() {
	fmt.Println("lib1.Lib1Func 调用 " + lib2.Lib2Func())
	// lib2.lib2Func()
}

func init() {
	fmt.Println("lib1.init ...")
}
