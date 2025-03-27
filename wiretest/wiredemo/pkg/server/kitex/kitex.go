package kitex

import (
	"context"
	kitex "github.com/cloudwego/kitex/server"
	"log/slog"
	"time"
)

type Server struct {
	//*gin.Engine
	//httpSrv *http.Server
	kitex.Server
	host   string
	port   int
	logger *slog.Logger
}

type Option func(s *Server)

func NewServer(server kitex.Server, logger *slog.Logger, opts ...Option) *Server {
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
		s.logger.Info("Server running %s \n", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.Server.Stop()
	if err != nil {
		s.logger.Info("Server forced to shutdown: ", err)
	}
	s.logger.Info("Server exiting")
	return nil
}
