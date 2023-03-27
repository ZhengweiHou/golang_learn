package main

import (
	"fmt"
)

// const定义常量（只读属性）

// const 可以用来定义枚举
const (
	// const()中可以使用关键字iota，常量计数器，iota会在每一行累加1
	// iota 只能用于const()中
	A = iota // iota计数从0开始
	B = iota
	C = iota

	D = 10 * iota
	E // 自动按上行规则补全 等价于 E = 10 * iota
	F // 等价于 E = 10 * iota

	G = iota + 10*iota // iota=6 多次使用不会累加，只和行数有关
	H                  // iota=7 7+7*10=77
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func main() {
	fmt.Println("A=", A) // 0
	fmt.Println("B=", B) // 1
	fmt.Println("C=", C) // 2
	fmt.Println("D=", D) // 30
	fmt.Println("E=", E) // 40
	fmt.Println("F=", F) // 50
	fmt.Println("G=", G) // 66
	fmt.Println("H=", H) // 77

	fmt.Println("a=", a) // 1
	fmt.Println("b=", b) // 2
	fmt.Println("c=", c) // 2
	fmt.Println("d=", d) // 3
	fmt.Println("e=", e) // 3
	fmt.Println("f=", f) // 4

	const hzw int = 100
	// hzw = 1000 // const定义的常量为只读属性，不可修改

}
