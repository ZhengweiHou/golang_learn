package main

import (
	"fmt"
	"testing"
)

func TestChannel1(t *testing.T) {
	c1 := make(chan string, 3)

	c1 <- "hhh"
	c1 <- "zzz"
	c1 <- "www"

	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)

	c1 <- "111"

	fmt.Println(len(c1)) // 1
	_, ok := <-c1
	fmt.Println(ok)      // true
	fmt.Println(len(c1)) // 0

	c1 <- "222"

	close(c1) // chanl中有数据，close关不掉

	fmt.Println(len(c1)) // 1
	_, ok = <-c1
	fmt.Println(ok)      // false
	fmt.Println(len(c1)) // 0

}
