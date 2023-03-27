package lib2

import (
	"fmt"
)

// 大写字母开头,称为导出，的标识符对象可以被外部包代码使用 类似于public
func Lib2Func() (r1 string) {
	fmt.Println("lib2.Lib2Func ...")
	lib2Func() // 未导出的方法或变量可以在包内部使用
	return "Lib2Func"
}

// 小写字母开头的标识符对象对包外不可见，但在整个包的内部是可见并可用 类似于private
func lib2Func() {
	fmt.Println("lib2.lib2Func ...")
}

func init() {
	fmt.Println("lib2.init ...")
}
