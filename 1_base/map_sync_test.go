package base

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncMap1(t *testing.T) {

	var m1 sync.Map
	for i := 0; i < 10; i++ {
		m1.Store(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	m1.Range(func(k, v any) bool {
		deled := m1.CompareAndDelete(k, v)
		fmt.Printf("deled:%t time:%s\n", deled, time.Now())
		return true
	})
}

func TestSyncMap2(t *testing.T) {

	var m1 sync.Map
	ob, ok := m1.Swap("1", "1111")
	fmt.Printf("ob:%v, ok:%t\n", ob, ok) // nil(原始值为nil) false(新增，没有替换)

	ob, ok = m1.Swap("1", "1111")
	fmt.Printf("ob:%v, ok:%t\n", ob, ok)
}
