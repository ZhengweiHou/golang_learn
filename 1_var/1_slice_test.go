package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	myarr := []int{1}
	// 追加一个元素
	myarr = append(myarr, 2)
	// 追加多个元素
	myarr = append(myarr, 3, 4)
	// 追加一个切片, ... 表示解包，不能省略
	myarr = append(myarr, []int{7, 8}...)
	// 在第一个位置插入元素
	myarr = append([]int{0}, myarr...)
	// 在中间插入一个切片(两个元素)
	myarr = append(myarr[:5], append([]int{5, 6}, myarr[5:]...)...)
	fmt.Println(myarr)
	fmt.Println(myarr[1:len(myarr)])
}

func Test2(t *testing.T) {

	myarr := []int{1}
	myarr = append(myarr, 2)

	arr2 := myarr
	arr2 = append(arr2, 3) // 不会修改原切片内容
	arr2[0] = 9            // 不会修改原切片内容
	fmt.Printf("%T  %v\n", myarr, myarr)
	fmt.Printf("%T  %v\n", arr2, arr2)

}
