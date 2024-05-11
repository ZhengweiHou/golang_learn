package base

import (
	"bytes"
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
	arr3 := &myarr
	(*arr3)[0] = 10 // 会修改原切片的数据
	fmt.Printf("%T  %v\n", myarr, myarr)
}

func TestSliceMapRange(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	// 迭代切片
	for i, v := range slice {
		fmt.Printf("slice[%d] = %d\n", i, v)
	}

	/*
		map 内部的实现是使用哈希表，哈希表的特性是无序的。所以，在使用 range 迭代一个 map 时，迭代顺序是不确定的
	*/
	// 迭代 map
	for k, v := range m {
		fmt.Printf("m[%s] = %d\n", k, v)
	}

}

func TestSliceRange2(t *testing.T) {

	type _tmp struct {
		i int
	}

	slice := []int{1, 2, 3, 4, 5}
	slice2 := []*_tmp{}
	// 迭代切片
	for _, v := range slice {
		slice2 = append(slice2, &_tmp{i: v})
	}
	for _, v2 := range slice2 {
		fmt.Printf("%v\n", v2)
	}

}
func TestSlice3(t *testing.T) {

	slice := []int{1, 2, 3, 4, 5}

	s1 := slice[2:]
	s2 := slice[:0]

	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)

	bt := []byte{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", bt)

	buff := &bytes.Buffer{}
	buff.Write(bt)
	fmt.Printf("%v\n", buff.Bytes())

	buff.Reset()
	fmt.Printf("%v\n", buff.Bytes())

}

func TestSli4(t *testing.T) {
	bt := []byte{1, 2, 3, 4, 5, 6}

	buff := &bytes.Buffer{}

	buff.Write(bt)
	sli1 := buff.Bytes()
	fmt.Printf("%v\n", sli1)

	buff.Reset()
	bt2 := []byte{9, 8, 7, 6}
	buff.Write(bt2) // 会覆盖上方的sli1
	sli2 := buff.Bytes()
	fmt.Printf("%v\n", sli2)
	fmt.Printf("%v\n", sli1)
}
