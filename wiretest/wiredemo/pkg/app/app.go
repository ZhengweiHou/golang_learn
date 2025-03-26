package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "wiredemo/pkg/log"
	"wiredemo/pkg/server"
)

type App struct {
	name    string
	servers []server.Server
	// Logger  *log.Logger
	Logger *slog.Logger
}

type Option func(a *App)

func NewApp(opts ...Option) (*App, func()) {
	a := &App{}
	for _, opt := range opts {
		opt(a)
	}
	cleanup := func() {
		fmt.Printf("wire cleanup!!!!!!")
		time.Sleep(time.Second)
	}
	return a, cleanup
}

func WithServer(servers ...server.Server) Option {
	return func(a *App) {
		a.servers = servers
	}
}

func WithName(name string) Option {
	return func(a *App) {
		a.name = name
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(a *App) {
		a.Logger = logger
	}
}
func (a *App) Run(ctx context.Context) error {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for _, srv := range a.servers {
		go func(srv server.Server) {
			err := srv.Start(ctx)
			if err != nil {
				a.Logger.Info(fmt.Sprintf("Server start err: %v", err))
			}
		}(srv)
	}

	select {
	case <-signals:
		// Received termination signal
		a.Logger.Info("Received termination signal")
	case <-ctx.Done():
		// Context canceled
		//log.Println("Context canceled")
		a.Logger.Info("Context canceled")
	}

	// Gracefully stop the servers
	for _, srv := range a.servers {
		err := srv.Stop(ctx)
		if err != nil {
			//log.Printf("Server stop err: %v", err)
			a.Logger.Info(fmt.Sprintf("Server stop err: %v", err))
		}
	}

	return nil
}
