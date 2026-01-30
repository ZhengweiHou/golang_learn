package servertest

import (
	"context"
	"log/slog"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/gzip"
)

func TestHelloHertzServer(t *testing.T) {
	h := server.Default(
		server.WithHostPorts(":8888"),
	)

	// 设置中间件
	h.Use(
		// Compressiong
		gzip.Gzip(gzip.DefaultCompression),
	)

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		slog.Info("=====")
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.Spin()
}
