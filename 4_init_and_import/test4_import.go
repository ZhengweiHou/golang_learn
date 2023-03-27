package main

import (
	"fmt"
	"golang_learn/4_init_and_import/lib1"
	_ "golang_learn/4_init_and_import/lib2" // 匿名import，不使用时不会报错
	mylib3 "golang_learn/4_init_and_import/lib3"
	"log"
)

func main() {
	log.Default().Fatalln("hahaha")
	fmt.Println("hello")
	lib1.Lib1Test()
	mylib3.Lib3func()
}
