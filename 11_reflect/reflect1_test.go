package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	type1 := reflect.TypeOf("123")
	value1 := reflect.ValueOf("123")

	fmt.Println(type1, value1)
	fmt.Println(type1.Kind())

	var c interface{}
	c = "hzw"
	fmt.Println(c)
	// c.(type1.Kind())
}

func Test2(t *testing.T) {
	// 创建一个实现接口的结构体实例
	myStructInstance := MyStruct{
		Id:   1,
		Name: "hzw"}

	// 将结构体实例转换为接口类型
	myInterfaceInstance := interface{}(&myStructInstance)

	// 使用反射获取接口的类型
	interfaceType := reflect.TypeOf(myInterfaceInstance)

	// 使用反射创建结构体实例的指针
	structPtr := reflect.New(interfaceType.Elem())

	// 使用反射将接口实例转换为结构体实例
	structPtr.Elem().Set(reflect.ValueOf(&myStructInstance).Elem())

	// 使用反射调用结构体的方法
	method := structPtr.MethodByName("MyMethod")
	method.Call(nil)

	// 将反射对象还原为原始结构体类型
	// originalStruct := structPtr.Interface().(MyStruct)

	// 将反射对象还原为原始结构体类型（通过 Elem() 获取指针指向的值）
	originalStruct := structPtr.Elem().Interface().(MyStruct)

	// 直接访问结构体的字段
	fmt.Println("Name:", originalStruct.Name)

}

func Test3(t *testing.T) {
	// 创建一个实现接口的结构体实例
	myStructInstance := &MyStruct{Id: 1, Name: "hzw"}

	// 将结构体实例转换为接口类型
	var ms MyInterface = myStructInstance

	// 使用反射获取接口的值
	value := reflect.ValueOf(ms)

	// 确保接口持有的是指针类型
	if value.Kind() == reflect.Ptr {
		// 获取指针指向的结构体值
		structValue := value.Elem()

		// 将反射对象还原为原始结构体类型
		originalStruct := structValue.Interface().(MyStruct)

		// 直接访问结构体的字段
		fmt.Println("Name:", originalStruct.Name)
	} else {
		fmt.Println("Not a pointer to a struct")
	}

}

type MyInterface interface {
	MyMethod()
}

type MyStruct struct {
	Id   int
	Name string
}

func (s MyStruct) MyMethod() {
	fmt.Println("Hello from MyMethod")
}
