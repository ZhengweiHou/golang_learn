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
