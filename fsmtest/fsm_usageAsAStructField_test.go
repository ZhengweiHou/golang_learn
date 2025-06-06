package fsmtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/looplab/fsm"
)

// Usage as a struct field
type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"enter_state": func(context context.Context, e *fsm.Event) {
				d.enterState(e)
			},
		},
	)

	return d
}

func (d *Door) enterState(e *fsm.Event) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}
func TestFsmAsStructField(t *testing.T) {
	door := NewDoor("heaven")

	err := door.FSM.Event(context.Background(), "open")
	if err != nil {
		fmt.Println(err)
	}

	err = door.FSM.Event(context.Background(), "close")
	if err != nil {
		fmt.Println(err)
	}
}
