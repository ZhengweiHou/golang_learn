//go:build wireinject
// +build wireinject

/*
 */

package wire

import (
	"log/slog"
	"wiredemo/internal/adapter/adapterhttp"
	"wiredemo/internal/adapter/adapterkitex"
	"wiredemo/internal/repository"
	"wiredemo/internal/server/http"
	"wiredemo/internal/server/kitex"
	"wiredemo/internal/service"
	"wiredemo/pkg/app"
	"wiredemo/pkg/log"
	http2 "wiredemo/pkg/server/http"
	kitex2 "wiredemo/pkg/server/kitex"

	//"aic.com/pkg/aicdb"
	"aic.com/pkg/aicgormdb"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// 应用服务器实现
var HttpServerSet = wire.NewSet(
	// http服务
	http.NewHTTPServer,
)

// KitexServerSet Kitex服务相关provider
var KitexServerSet = wire.NewSet(
	kitex.NewKitexOriginalServer,
	kitex.NewKitexServer,
	kitex2.NewHzwKCReporter,
)

// http 处理器
var AdapterhttpSet = wire.NewSet(
	adapterhttp.NewHzwController,
	adapterhttp.NewHzw2Controller,
)

// AsapterKitexSet kitex处理器
var AdapterkitexSet = wire.NewSet(
	adapterkitex.NewHzwKitexCtl,
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
	kitexServer *kitex2.Server,
	// grpcServer *grpc.Server,
	logger *slog.Logger,
) (*app.App, func()) {
	return app.NewApp(
		app.WithServer(
			httpServer,
			kitexServer,
		),
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
		HttpServerSet,
		AdapterhttpSet,
		KitexServerSet,
		AdapterkitexSet,
		ServiceSet,
		newApp,
	))
}
