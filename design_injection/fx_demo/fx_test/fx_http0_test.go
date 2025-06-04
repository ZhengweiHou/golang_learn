package fx_test

import (
	"context"
	"net/http"
	"testing"

	"go.uber.org/fx"
)

func TestHttp0(t *testing.T) {
	app := fx.New(
		fx.Provide(
			NewHTTP0ServerWithHook,
		),
		fx.Invoke(func(*http.Server) {}),
	)
	app.Run()
}

func NewHTTP0ServerWithHook(lc fx.Lifecycle) *http.Server {
	srv := &http.Server{Addr: ":8888"}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {

			go srv.ListenAndServe()
			// go srv.Serve(l net.Listener)
			return nil
		},
		OnStop: func(context.Context) error {
			return srv.Shutdown(context.Background())
		},
	})
	return srv
}
