package main

import "fmt"

func main() {
	// 接口测试
	var pI MyInterface1
	pI = &Person1{age: 20, name: "李四"}

	var p Person1
	p = Person1{age: 10, name: "张三"}

	pI.SayHello()
	p.SayHello()

	changeName1(&pI)
	changeName2(&p)

	pI.SayHello()
	p.SayHello()
}

func changeName1(p *MyInterface1) {
	// p
}
func changeName2(p *Person1) {
	p.name = "hehe"
}

// 定义一个接口
type MyInterface1 interface {
	SayHello()
	GetAge() int
	Cname(string)
}

// 定义类
type Person1 struct {
	age  int
	name string
}

func (p *Person1) SayHello() {
	fmt.Println("hello ", p.name)
}

func (p *Person1) GetAge() int {
	return p.age
}

func (p *Person1) Cname(name string) {
	p.name = name
}
