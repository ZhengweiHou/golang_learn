package wirehello

import (
	"fmt"
)

// Service 定义一个服务接口
type Service interface {
	DoSomething()
}

// ServiceA 实现 Service 接口
type ServiceA struct {
	Msg string
}

func (s *ServiceA) DoSomething() {
	fmt.Println("ServiceA is doing something" + s.Msg)
}

// ServiceB 实现 Service 接口
type ServiceB struct{}

func (s *ServiceB) DoSomething() {
	fmt.Println("ServiceB is doing something")
}

// App 定义一个应用结构体，依赖服务切片
type App struct {
	Services []Service
}

// NewApp 创建一个新的 App 实例
func NewApp(services []Service) *App {
	return &App{
		Services: services,
	}
}
func NewAppA(services *ServiceA) *App {
	return &App{}
}

// Run 运行应用
func (a *App) Run() {
	for _, service := range a.Services {
		service.DoSomething()
	}
}
