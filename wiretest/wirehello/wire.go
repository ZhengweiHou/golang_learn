/*
//go:build wireinject
// +build wireinject
*/

// 标示 该文件只在运行wire是被编译

package wirehello

import "github.com/google/wire"

// InitializeEvent 声明injector的函数签名
func InitializeEvent(msg string) Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{} //返回值没有实际意义，只需符合函数签名即可
}

var wireSet = wire.NewSet(
	//	wire.Struct(new(App), "*"),
	//	wire.Bind(new(Service), new(*ServiceA)),
	//	wire.Bind(new(Service), new(*ServiceB)),
	wire.Value(&ServiceA{}),
	wire.Value(&ServiceA{"2"}),
	wire.Value(&ServiceB{}),
	NewAppA,
	// wire.Value(&App{}),
)

func InitApp() *App {
	panic(wire.Build(wireSet))
}
