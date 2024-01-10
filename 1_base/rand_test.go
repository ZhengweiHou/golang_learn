package base

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand1(t *testing.T) {
	fmt.Println(rand.Int31())
	fmt.Println(rand.Int31n(100))
	fmt.Println(rand.Int63())
	fmt.Println(rand.Int63n(100))
}

func TestRand2(t *testing.T) {
	// 设置种子，每次生成不一样的值
	rand.Seed((time.Now().UnixNano()))
	fmt.Println(rand.Intn(100))
	rand.Seed((time.Now().UnixNano()))
	fmt.Println(rand.Intn(100))
}

func TestRand3(t *testing.T) {
	// 设置种子，Unix默认是按秒为维度，下面大概率随机数一样
	rand.Seed((time.Now().Unix()))
	fmt.Println(rand.Int31n(100))
	rand.Seed((time.Now().Unix()))
	fmt.Println(rand.Int31n(100))
}
