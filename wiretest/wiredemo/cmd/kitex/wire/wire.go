//go:build wireinject
// +build wireinject

package wire

import (
	"aic.com/pkg/aicgormdb"
	"github.com/google/wire"
	"github.com/spf13/viper"

	"log/slog"
	"wiredemo/internal/controller"
	"wiredemo/internal/repository"
	"wiredemo/internal/server"
	"wiredemo/internal/service"
	"wiredemo/pkg/app"
	"wiredemo/pkg/log"
	"wiredemo/pkg/server/kitex"
)

// 应用服务器实现
var ServerSet = wire.NewSet(
	// kitex服务
	server.NewKitexServer,
)

// http 处理器
var ControllerSet = wire.NewSet(
	controller.NewHzwController,
	controller.NewHelloController,
	controller.NewBybyController,
)

// 业务服务
var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewHzwService,
	//wire.Bind(new(service.IHzwService), new(*service.HzwService)), // 若实现构造器函数签名返回的不是接口，则需要绑定映射关系，否则wire无法通过接口关联实现依赖
)

// build App
func newApp(
	kitexServer *kitex.Server,
	logger *slog.Logger,
	// grpcServer *grpc.Server,
) (*app.App, func()) {
	return app.NewApp(
		app.WithServer(kitexServer),
		app.WithName("wiredemo-kitex-server"),
		app.WithLogger(logger),
	)
}

// wire 整合构建
func NewWire(*viper.Viper) (*app.App, func(), error) {
	panic(wire.Build(
		log.LogWireSet,
		aicgormdb.DbWireSet, // db子包中定义的wireset
		repository.RepositoryWireSet,

		ServerSet,
		ControllerSet,
		ServiceSet,
		newApp,
	))
}
