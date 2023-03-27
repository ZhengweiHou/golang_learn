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
}
