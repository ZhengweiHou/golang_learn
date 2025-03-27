package kitex

import (
	"context"
	kitex "github.com/cloudwego/kitex/server"
	"time"
	"wiredemo/pkg/log"
)

type Server struct {
	//*gin.Engine
	//httpSrv *http.Server
	kitex.Server
	host   string
	port   int
	logger *log.Logger
}

type Option func(s *Server)

func NewServer(server kitex.Server, logger *log.Logger, opts ...Option) *Server {
	s := &Server{
		Server: server,
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithServerHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithServerPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func (s *Server) Start(ctx context.Context) error {
	err := s.Server.Run()

	if err != nil {
		s.logger.Sugar().Fatalf("Server running %s \n", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Sugar().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.Server.Stop()
	if err != nil {
		s.logger.Sugar().Fatal("Server forced to shutdown: ", err)
	}
	s.logger.Sugar().Info("Server exiting")
	return nil
}
