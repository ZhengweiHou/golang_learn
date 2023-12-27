package fsmtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/looplab/fsm"
)

// basic examle
func TestFsm1(t *testing.T) {
	fsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(fsm.Current())

	err := fsm.Event(context.Background(), "open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	err = fsm.Event(context.Background(), "close")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())
}
