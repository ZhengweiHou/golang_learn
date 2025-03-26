//go:build wireinject
// +build wireinject

/*
 */

package wire

import (
	"log/slog"
	"wiredemo/internal/controller"
	"wiredemo/internal/repository"
	"wiredemo/internal/server/http"
	"wiredemo/internal/service"
	"wiredemo/pkg/app"
	"wiredemo/pkg/log"
	http2 "wiredemo/pkg/server/http"

	//"aic.com/pkg/aicdb"
	"aic.com/pkg/aicgormdb"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// 应用服务器实现
var ServerSet = wire.NewSet(
	// http服务
	http.NewHTTPServer,
)

// http 处理器
var ControllerSet = wire.NewSet(
	controller.NewHzwController,
	controller.NewHzw2Controller,
)

// 业务服务
var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewHzwService,
	service.NewHzw2Service,
	//wire.Bind(new(service.IHzwService), new(*service.HzwService)), // 若实现构造器函数签名返回的不是接口，则需要绑定映射关系，否则wire无法通过接口关联实现依赖
)

// build App
func newApp(
	httpServer *http2.Server,
	// grpcServer *grpc.Server,
	logger *slog.Logger,
) (*app.App, func()) {
	return app.NewApp(
		app.WithServer(httpServer),
		app.WithName("wiredemo-server"),
		app.WithLogger(logger),
	)
}

// wire 整合构建
// func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
func NewWire(*viper.Viper) (*app.App, func(), error) {
	panic(wire.Build(
		log.LogWireSet,
		aicgormdb.DbWireSet,
		repository.RepositoryWireSet,
		//aicdb.DbWireSet,
		ServerSet,
		ControllerSet,
		ServiceSet,
		newApp,
	))
}
