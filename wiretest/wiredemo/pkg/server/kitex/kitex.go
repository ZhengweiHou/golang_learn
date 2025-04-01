package kitex

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/cloudwego/kitex/pkg/remote/connpool"
	"github.com/cloudwego/kitex/server"
)

type Server struct {
	ksvr   server.Server
	logger *slog.Logger
}
type Option func(s *Server)

// NewServer (engine *gin.Engine, logger *log.Logger, opts ...Option) *Server {
func NewServer(svr server.Server, logger *slog.Logger, opts ...Option) *Server {
	s := &Server{
		ksvr:   svr,
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) Start(ctx context.Context) error {
	sinfos := s.ksvr.GetServiceInfos()
	// 格式化打印服务信息
	for sname, sinfo := range sinfos {
		s.logger.Info("Kitex服务信息",
			"服务名称", sname,
			"方法数量", len(sinfo.Methods),
			"方法列表", sinfo.Methods,
			"额外信息", sinfo.Extra,
		)
	}

	s.logger.Info("Kitex服务启动",
		"总服务数", len(sinfos),
		"详细信息", sinfos,
	)
	return s.ksvr.Run()
}
func (s *Server) Stop(ctx context.Context) error {
	return s.ksvr.Stop()
}

// HzwKCReporter 自定义连接池监控
type HzwKCReporter struct{}

func NewHzwKCReporter() *HzwKCReporter {
	return &HzwKCReporter{}
}

func (r *HzwKCReporter) ConnSucceed(poolType connpool.ConnectionPoolType, serviceName string, addr net.Addr) {
	fmt.Printf("ConnSucceed poolType:%d, serviceName:%s, addr:%v\n", poolType, serviceName, addr)
}
func (r *HzwKCReporter) ConnFailed(poolType connpool.ConnectionPoolType, serviceName string, addr net.Addr) {
	fmt.Printf("ConnFailed poolType:%d, serviceName:%s, addr:%v\n", poolType, serviceName, addr)
}
func (r *HzwKCReporter) ReuseSucceed(poolType connpool.ConnectionPoolType, serviceName string, addr net.Addr) {
	fmt.Printf("ReuseSucceed poolType:%d, serviceName:%s, addr:%v\n", poolType, serviceName, addr)
}
