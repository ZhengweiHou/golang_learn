//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"wire-hzw/internal/handler"
	"wire-hzw/internal/repository"
	"wire-hzw/internal/server"
	"wire-hzw/internal/service"
	"wire-hzw/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(
	repository.NewDb,
	repository.NewRepository,
	repository.NewUserRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

func NewWire(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
