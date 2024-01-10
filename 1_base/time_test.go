package base

import (
	"fmt"
	"testing"
	"time"
)

func TestTime1(t *testing.T) {
	var i *int
	i2 := 1000
	i = &i2
	var du time.Duration
	du = time.Millisecond * time.Duration(*i)
	fmt.Println(time.Now())
	time.Sleep(du)
	fmt.Println(time.Now())
}
