package main

import (
	"context"
	"log"
	"os"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

/*
Fx 依赖注入框架示例

本示例展示Fx框架的核心功能：
1. 依赖注入
2. 生命周期管理
3. 模块化设计
4. 可选依赖
*/

// 1. 基本服务定义
// Greeter 服务接口定义
type Greeter interface {
	Greet() string
}

// greeter 服务实现
type greeter struct {
	name string
}

// Greet 实现接口方法
func (g *greeter) Greet() string {
	return "Hello, " + g.name
}

// NewGreeter 是Greeter的构造函数
// Fx会调用此函数创建Greeter实例
// 注意：构造函数通常返回接口类型，而不是具体实现
func NewGreeter() Greeter {
	return &greeter{name: "Fx User"}
}

// 2. 带依赖的服务
// Logger 日志服务接口
type Logger interface {
	Log(msg string)
}

// logger 日志服务实现
type logger struct{}

// Log 实现日志记录
func (l *logger) Log(msg string) {
	log.Println(msg)
}

// NewLogger 日志服务构造函数
func NewLogger() Logger {
	return &logger{}
}

// GreeterService 组合服务，演示依赖注入
type GreeterService struct {
	fx.In   // 标记为需要依赖注入的结构体
	Greeter Greeter
	Logger  Logger
}

// Run 执行业务逻辑，自动注入的依赖可直接使用
func (s *GreeterService) Run() {
	msg := s.Greeter.Greet()
	s.Logger.Log(msg)
}

// 3. 生命周期管理
// registerLifecycleHooks 注册应用生命周期钩子
// lc 由Fx自动注入的生命周期管理器
// greeter 由Fx自动注入的Greeter服务
func registerLifecycleHooks(lc fx.Lifecycle, greeter Greeter) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// 应用启动时执行
			log.Println("Starting application...")
			log.Println(greeter.Greet())
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// 应用停止时执行
			log.Println("Stopping application...")
			return nil
		},
	})
}

// 4. 可选依赖
// OptionalService 演示可选依赖
type OptionalService struct {
	fx.In
	Greeter Greeter `optional:"true"` // 标记为可选依赖
	Logger  Logger
}

// Run 执行业务逻辑，处理可选依赖
func (s *OptionalService) Run() {
	if s.Greeter != nil {
		s.Logger.Log(s.Greeter.Greet())
	} else {
		// 可选依赖不存在时的处理
		s.Logger.Log("No greeter available")
	}
}

// 5. 模块化设计

// GreeterModule Greeter服务模块
// 使用fx.Module将相关组件组织在一起
var GreeterModule = fx.Module("greeter",
	fx.Provide(NewGreeter), // 注册服务提供者
)

// LoggerModule 日志服务模块
var LoggerModule = fx.Module("logger",
	fx.Provide(NewLogger),
)

func main() {
	// 创建Fx应用
	app := fx.New(
		// 注册模块
		GreeterModule,
		LoggerModule,

		// 注册生命周期钩子
		fx.Invoke(registerLifecycleHooks),

		// 运行主服务
		fx.Invoke(func(s GreeterService) {
			s.Run()
		}),

		// 演示可选依赖
		fx.Invoke(func(s OptionalService) {
			s.Run()
		}),

		// 配置选项：自定义Fx事件日志
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.ConsoleLogger{W: os.Stdout}
		}),
	)

	// 启动应用
	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	// 优雅关闭
	if err := app.Stop(context.Background()); err != nil {
		log.Fatal(err)
	}
}
