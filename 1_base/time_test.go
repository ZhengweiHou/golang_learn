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
func TestTime2(t *testing.T) {

	for i := 1; i < 3; i++ {
		fmt.Printf("%T,%d\n", i, i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}
