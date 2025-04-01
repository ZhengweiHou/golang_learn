# wiredemo 项目说明

## 项目概述
基于 Go 语言的微服务示例项目，使用 Wire 依赖注入工具、Kitex 微服务框架 和 gorm 数据库访问框架 构建。

## 详细目录结构说明
### 核心目录
```
wiredemo/
│
├── api/                 # API接口定义层
│   ├── grpc/            # gRPC协议接口
│   ├── http/            # HTTP REST接口
│   │   └── v1/          # API v1版本
│   │       ├── errors.go # 错误及业务异常定义
│   │       └── v1.go    # 基础http实例定义
│   └── kitex/           # Kitex微服务接口
│       ├── idl/         # Thrift IDL文件
│       │   └── hzw.thrift # 服务定义
│       └── hzw/         # Kitex服务实现(根据idl自动生成)
│           ├── hzw.go   # 服务主文件
│           └── hzwservice/ # 服务实现
│               ├── client.go # 客户端代码
│               └── server.go # 服务端代码
│
├── cmd/                 # 应用入口点
│   └── server/          # 服务主程序
│       ├── main.go      # 程序入口
│       └── wire/        # 依赖注入
│           ├── wire.go  # 依赖定义
│           └── wire_gen.go # 生成的依赖代码
│
├── config/              # 配置管理
│   └── app.yml          # 全局应用配置
│
├── docs/                # 项目文档（Swagger）
│   ├── docs.go          # Go文档
│   ├── swagger.json     # Swagger JSON
│   └── swagger.yaml     # Swagger YAML
│
├── internal/            # 内部实现
│   ├── adapter/         # 协议适配层
│   │   ├── adapterhttp/ # HTTP适配器
│   │   │   └── hzwcontroller.go
│   │   └── adapterkitex/ # Kitex适配器
│   │       └── handler_HzwService.go
│   ├── repository/      # 数据仓库层
│   │   ├── idl/         # 数据访idl定义
│   │   │   └── xxx.yaml # 基础实体定义
│   │   ├── dao/         # 数据访问对象(由idl生成)
│   │   │   └── xxxdao.go
│   │   ├── model/       # 数据模型(由idl生成)
│   │   │   └── xxx.go
│   │   └── wire.go      # 仓库层依赖注入（TODO 后续也自动生成）
│   ├── server/          # 服务层
│   │   ├── http/        # HTTP服务
│   │   │   └── http.go  # HTTP服务实现(用于加载adapter中的各处理器)
│   │   └── kitex/       # Kitex服务
│   │       └── kitex.go # Kitex服务实现(用于加载adapter中的各处理器)
│   └── service/         # 业务服务层（核心业务逻辑在该层实现，由协议适配层接入调用）
│       ├── baseservice.go # 服务基类，事务管理实现注入
│       ├── hzw2service.go # 业务服务1
│       └── hzwservice.go  # 业务服务2
│
└── test/                # 测试代码
    └── xxxx_test.go # 测试代码
```

## 架构分层

1. **API层**：定义外部接口协议
   - HTTP/RESTful API
   - gRPC API
   - Kitex微服务API

2. **Service层**：实现核心业务逻辑
   - 业务服务实现
   - 事务管理
   - 业务规则验证

3. **Repository层**：处理数据访问
   - 数据库操作
   - 缓存访问
   - 数据转换

4. **Model层**：定义数据结构
   - 数据库模型
   - DTO对象
   - 请求/响应结构

5. **Adapter层**：协议转换适配
   - HTTP请求处理
   - gRPC消息转换
   - 协议适配实现

## 快速开始

```bash
....
```

## 依赖工具

- Go 1.18+
- Wire (依赖注入)
- Kitex (微服务框架)
- GORM (ORM框架)
- Swagger (API文档)
- 
