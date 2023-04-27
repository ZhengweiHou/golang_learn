package main

import "fmt"

func main() {
	// 接口测试
	var empty any // any 实际上就是 空接口 interface{}
	empty = Person{age: 20, name: "empty"}
	if realType, ok := empty.(Person); ok {
		fmt.Println("type is Person ", realType)
	} else if realType, ok := empty.(Animal); ok {
		fmt.Println("type is Animal", realType)
	}

	// 继承测试
	stu := Student{school: "middle"}
	stu.name = "leo"
	stu.age = 30
	fmt.Print(stu.name)
	stu.SayHello()

	

}

// 定义一个接口
type MyInterface interface {
	SayHello()
	GetAge() int
}

// 定义类
type Animal struct{}
type Person struct {
	age  int
	name string
}

func (p Person) SayHello() {
	fmt.Println("hello ", p.name)
}

func (p Person) GetAge() int {
	return p.age
}

// 继承
type Student struct {
	Person
	school string
}
