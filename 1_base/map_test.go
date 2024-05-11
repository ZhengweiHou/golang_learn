package base

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestMap1(t *testing.T) {
	m := make(map[int]int)

	fmt.Printf("len:%v\n", len(m))
	for i := 0; i < 60; i++ {
		n := rand.Int()
		m[i] = n
	}

	fmt.Printf("len:%v\n", len(m))
	for k := range m {
		delete(m, k)
	}

	fmt.Printf("len:%v\n", len(m))
}
func TestMap2(t *testing.T) {
	var wg sync.WaitGroup
	var rmap = make(map[string]interface{})
	rmap["k1"] = "v1"
	rmap["t2"] = time.Now()

	wg.Add(1)
	go func(rmap2 map[string]interface{}) {
		rmap2["k3"] = "v3"   // 1️⃣ map为指针传递，此处修改会影响外数据
		rmap2["k1"] = "v1_2" // 1️⃣ map为指针传递，此处修改会影响外数据

		v := rmap2["k1"] // 2️⃣ 此处是值复制
		wg.Wait()
		fmt.Println(v)
	}(rmap)

	time.Sleep(time.Microsecond)
	rb, _ := json.MarshalIndent(rmap, "", "    ")
	fmt.Printf("%s\n", rb) // 1️⃣ 指针传递修改会被修改

	rmap["k1"] = "v1_3" // 2️⃣ 值复制处不会被影响
	wg.Done()
	time.Sleep(time.Microsecond)
}

func TestMap3(t *testing.T) {
	var rmap = make(map[string]interface{})
	rmap["k1"] = "v1"

	func(rmap2 map[string]interface{}) {
		// vp := rmap2["k1"] // 此处是值复制
		fmt.Printf("%p\n", &rmap2)
	}(rmap)

}

func TestMap4(t *testing.T) {
	var map1 map[int32]*string

	if map1 == nil {
		fmt.Println("map1 is nil")
	}
}
