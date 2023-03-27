package main

import (
	"fmt"
	"math"
	"unsafe"
)

// 有符号整形
func Integer() {
	var num8 int8 = -127
	var num16 int16 = 32767
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	var num int = math.MaxInt

	fmt.Printf("num8的类型:%T ,num8的大小:%d ,num8是:%d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型:%T ,num16的大小:%d ,num16是:%d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型:%T ,num32的大小:%d ,num32是:%d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型:%T ,num64的大小:%d ,num64是:%d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型:%T ,num的大小:%d ,num是:%d\n", num, unsafe.Sizeof(num), num)

}

//无符号整形
func UInteger() {
	var num8 uint8 = 127
	//var num8 uint8 = -127 // 无符号整形无法赋值负数
	var num16 uint16 = 32767
	var num32 uint32 = math.MaxInt32
	var num64 uint64 = math.MaxInt64
	var num uint = math.MaxInt

	fmt.Printf("num8的类型:%T ,num8的大小:%d ,num8是:%d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型:%T ,num16的大小:%d ,num16是:%d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型:%T ,num32的大小:%d ,num32是:%d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型:%T ,num64的大小:%d ,num64是:%d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型:%T ,num的大小:%d ,num是:%d\n", num, unsafe.Sizeof(num), num)

}
func main() {
	Integer()
    fmt.Println("=======")
	UInteger()
}
