# Fx 依赖注入框架示例详解

本示例展示了 Go 语言中 Uber 的 Fx 依赖注入框架的核心功能。

## 1. Fx 框架简介

Fx 是一个用于 Go 的依赖注入框架，主要功能包括：
- 依赖注入管理
- 生命周期控制
- 模块化设计
- 易于测试

## 2. 示例代码解析

### 2.1 基本服务定义

```go
type Greeter interface {
    Greet() string
}

type greeter struct {
    name string
}

func NewGreeter() Greeter {
    return &greeter{name: "Fx User"}
}
```

- 定义服务接口和实现
- 通过`fx.Provide`注册服务提供者
- Fx 会自动管理服务的创建和依赖

### 2.2 依赖注入

```go
type GreeterService struct {
    fx.In
    Greeter Greeter
    Logger  Logger
}
```

- 使用`fx.In`结构体标签实现自动注入
- Fx 会自动解析并注入所需的依赖项
- 依赖关系在应用启动时自动解决

### 2.3 生命周期管理

```go
func registerLifecycleHooks(lc fx.Lifecycle, greeter Greeter) {
    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            log.Println("Starting application...")
            return nil
        },
        OnStop: func(ctx context.Context) error {
            log.Println("Stopping application...")
            return nil
        },
    })
}
```

- 通过`fx.Lifecycle`注册生命周期钩子
- `OnStart`: 应用启动时执行
- `OnStop`: 应用停止时执行
- 支持优雅关闭

### 2.4 可选依赖

```go
type OptionalService struct {
    fx.In
    Greeter Greeter `optional:"true"`
    Logger  Logger
}
```

- 使用`optional:"true"`标记可选依赖
- 如果依赖不可用，不会报错而是设为nil
- 需要在代码中检查nil情况

### 2.5 模块化设计

```go
var GreeterModule = fx.Module("greeter",
    fx.Provide(NewGreeter),
)

var LoggerModule = fx.Module("logger",
    fx.Provide(NewLogger),
)
```

- 使用`fx.Module`组织相关组件
- 提高代码的可维护性和复用性
- 模块可以独立测试和重用

## 3. 运行示例

1. 确保已安装Go 1.16+
2. 运行命令：
   ```bash
   go run fx_demo.go
   ```
3. 预期输出：
   ```
   Starting application...
   Hello, Fx User
   Hello, Fx User
   No greeter available
   Stopping application...
   ```

## 4. 最佳实践

1. 使用接口定义服务契约
2. 保持构造函数简单
3. 合理使用模块组织代码
4. 善用生命周期管理资源
5. 为关键服务添加健康检查
