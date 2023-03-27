package main

import "fmt"

/*
方法函数定义
func 函数名(参数列表)(返回值列表){
	函数执行体
}
*/

// 返回一个返回值
func foo1(a string, b int) int {
	fmt.Println("---- foo1 ----")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	return 100
}

// 返回多个返回值
func foo2(a string, b int) (int, string) {
	fmt.Println("---- foo2 ----")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	return 100, "def"
}

// 多个返回值，有形参名
func foo3(a int) (r1 int, r2 int) {
	fmt.Println("---- foo3 ----")
	fmt.Println("a=", a)
	fmt.Println("r1=", r1) // 返回值参数可以在方法体中使用，且有初始值0
	fmt.Println("r2=", r2)

	r2 = 100 // 给返回值形参赋值方式，返回值
	r1 = 200
	return
}

func foo4(a int) (r1, r2 int) {
	fmt.Println("---- foo4 ----")
	fmt.Println("a=", a)

	return 100, 200
}

func main() {
	c := foo1("abc", 100)
	fmt.Println(c)

	e, f := foo2("abc", 100) // 同时接收方法返回的多个值
	fmt.Println("e=", e, "f=", f)

	g, h := foo3(100) // 同时接收方法返回的多个值
	fmt.Println("g=", g, "h=", h)

	i, j := foo4(100) // 同时接收方法返回的多个值
	fmt.Println("i=", i, "j=", j)

	var k, l int = foo4(100)
	fmt.Println("k=", k, "l=", l)

	//var m, n int = foo2("abc", 100) // 返回值类型要对应，无法用int接收string返回值
	var m, n = foo2("abc", 100)
	fmt.Println("m=", m, "n=", n)
}
