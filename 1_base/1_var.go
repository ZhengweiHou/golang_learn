package base

import (
	"fmt"
	"log"
	"testing"
)

/*
	四种变量的声明方式
*/

// 全局变量只能使用1,2,3中声明方式
var ga int
var gb int = 100
var gc = 100

//gd := 100 // :=不能用来定义全局变量

func TestVar1(t *testing.T) {

	fmt.Println("hello")

	// 第一种方式
	var a int
	var aa string

	// 第二种方式 声明一个变量，并初始化一个值
	var b int = 100

	// 第三种方式 初始化的时候，去掉数据类型，Go语言通过值自动匹配类型
	var c = 100

	// 第四种方式 短声明
	d := 100
	e := true
	f := 3.14

	fmt.Printf("a  type: %-7T; value: %d\n", a, a)
	fmt.Printf("aa type: %-7T; value: %s\n", aa, aa)
	fmt.Printf("b  type: %-7T; value: %d\n", b, b)
	fmt.Printf("c  type: %-7T; value: %d\n", c, c)
	fmt.Printf("d  type: %-7T; value: %d\n", d, d)
	fmt.Printf("e  type: %-7T;", e)
	fmt.Println(" value:", e)
	fmt.Printf("f  type: %-7T; value: %f\n", f, f)

	// 定义多变量
	var g, h int
	var i, j = 100, "abcd"

	fmt.Printf("g  type: %-7T; value: %d\n", g, g)
	fmt.Printf("h  type: %-7T; value: %d\n", h, h)
	fmt.Printf("i  type: %-7T; value: %d\n", i, i)
	fmt.Printf("j  type: %-7T; value: %s\n", j, j)

	// 多行定义变量
	var (
		k int = 100
		l     = 100
	//	m := 100  // 不允许
	)
	fmt.Println("k=", k, " l=", l)

	log.Println("heheh")

	//使用表达式 new(Type) 将创建一个Type类型的匿名变量，初始化为Type类型的零值，然后返回变量地址，返回的指针类型为*Type
	ptr := new(int)
	fmt.Println("ptr address: ", ptr)
	fmt.Println("ptr value: ", *ptr) // * 后面接指针变量，表示从内存地址中取出值
}
