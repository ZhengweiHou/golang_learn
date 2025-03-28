package server

import (
	"fmt"
	"github.com/cloudwego/kitex/server"
	"github.com/spf13/viper"
	"log/slog"
	"net"
	"wiredemo/internal/controller"
	"wiredemo/pkg/kitex/kitex_gen/api/byby"
	"wiredemo/pkg/kitex/kitex_gen/api/hello"
	"wiredemo/pkg/server/kitex"
)

func NewKitexServer(
	logger *slog.Logger,
	conf *viper.Viper,
	helloc *controller.HelloController,
	bybyc *controller.BybyController,
) *kitex.Server {
	ip := fmt.Sprintf("%s:%d", conf.GetString("kitex.host"), conf.GetInt("kitex.port"))
	addr, err := net.ResolveTCPAddr("tcp", ip)
	if err != nil {
		logger.Error("%s", err)
		return nil
	}

	// 怎么加载多中server，怎么识别需要加载哪些controller
	srv := hello.NewServer(helloc, server.WithServiceAddr(addr))
	err = byby.RegisterService(srv, bybyc)

	if err != nil {
		logger.Error("%s", err)
		return nil
	}

	s := kitex.NewServer(
		srv,
		logger,
	)

	return s
}
