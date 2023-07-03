package main

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var a *string
	fmt.Print(a)
	if a != nil {
		fmt.Print(a)
	}
}
