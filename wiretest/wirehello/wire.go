//go:build wireinject
// +build wireinject

// 标示 该文件只在运行wire是被编译

package wirehello

import "github.com/google/wire"

// InitializeEvent 声明injector的函数签名
func InitializeEvent(msg string) Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{} //返回值没有实际意义，只需符合函数签名即可
}
