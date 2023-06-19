package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMap1(t *testing.T) {
	m := make(map[int]int)

	fmt.Printf("len:%v\n", len(m))
	for i := 0; i < 60; i++ {
		n := rand.Int()
		m[i] = n
	}

	fmt.Printf("len:%v\n", len(m))
	for k, _ := range m {
		delete(m, k)
	}

	fmt.Printf("len:%v\n", len(m))
}
