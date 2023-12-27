package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type MT struct {
	m sync.RWMutex
}

func Test_mutex1(t *testing.T) {
	var m1 sync.RWMutex
	// var m2 sync.RWMutex

	// mth := func(m *sync.RWMutex, msg string) {
	// 	for {
	// 		m.Lock()
	// 		fmt.Println(msg)
	// 		time.Sleep(time.Second * 1)
	// 		m.Unlock()

	// 	}
	// }
	mthR := func(m *sync.RWMutex, msg string) {
		for {
			m.RLock()
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
			m.RUnlock()

		}
	}

	go mthR(&m1, "Rm1g1")
	go mthR(&m1, "Rm1g2")
	// go mthR(&m2, "Rm2g1")
	// go mthR(&m2, "Rm2g2")

	time.Sleep(time.Second * 3)
	m1.Lock()
	time.Sleep(time.Second * 3)
	m1.Unlock()

	time.Sleep(time.Second * 5)
}

func Test_mutex2(t *testing.T) {
	// var m1 sync.RWMutex
	// var m2 sync.RWMutex
	m1 := &MT{}
	m2 := &MT{}

	mth := func(m *MT, msg string) {
		for {
			m.m.Lock()
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
			m.m.Unlock()

		}
	}

	go mth(m1, "m1g1")
	go mth(m1, "m1g2")
	go mth(m1, "m1g3")
	go mth(m2, "m2g1")
	go mth(m2, "m2g2")
	go mth(m2, "m2g3")

	time.Sleep(time.Second * 60)
}

// 协程内锁是否可重入
func Test_mutex3(t *testing.T) {
	m1 := &MT{}

	m1.m.RLock()
	m1.m.RUnlock()
	m1.m.RUnlock() // 不可以重复解锁
	// m1.m.Unlock()

	m1.m.TryRLock()

	m1.m.RLock()
	fmt.Println("in RLock")
	m1.m.Lock()
	fmt.Println("in Lock")
	m1.m.Unlock()
	fmt.Println("out Lock")
	m1.m.RUnlock()
	fmt.Println("out RLock")

}
