package hellowire

import "testing"

func TestWireHello(t *testing.T) {
	// 普通手动注入
	message := NewMessage("hello world")
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	event.Start()

	// wire生成的自动注入
	event2 := InitializeEvent("hello world2")
	event2.Start()
}
