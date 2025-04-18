package hello

import "testing"

func TestHzw(t *testing.T) {
	hctl := NewHController(NewHService())
	msg := hctl.HelloCtl("hello")
	t.Log(msg)
}
