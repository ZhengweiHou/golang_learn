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
	"github.com/hertz-contrib/pprof"
)

func TestDemoServerTest(t *testing.T) {
	h := server.Default(
		server.WithHostPorts(":8888"),
	)

	// pprof支持
	// pprof.Register(h)
	pprof.Register(h, "debug/pprof")

	// sort.Interface

	// 设置中间件
	h.Use(
		// Compressiong
		gzip.Gzip(gzip.DefaultCompression),
		// hzwhertzmw1,
		// hzwhertzmw2,
		// hzwhertzmw3,
	)

	// 路由
	slog.Info("=====")
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "root"})
	})
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.GET("/*path", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong any"})
	})

	h.GET("/v1/*path", func(ctx context.Context, c *app.RequestContext) {
		// body := c.Request.BodyBuffer()
		c.JSON(consts.StatusOK, utils.H{"message": "pong any"})
	})

	/* g1 和 g2 虽然 组path一样，但其中间件不同
	group是一个临时的，类似于builder一样，在未创建Handler时，无任何效果
	*/
	g1 := h.Group("/v1")
	g1.Use(hzwhertzmw1)
	g1.GET("/1", pongf("pong v1-1")).Use(hzwhertzmw2) // 此处handler后的mw不会生效！！
	g1.Use(hzwhertzmw3)                               // 此处mw不生效
	/* 中间件要在handler前设置，否则不会生效 */

	g2 := h.Group("/v1")
	g2.Use(hzwhertzmw2)
	g2.GET("/2", pongf("pong v1-2"))

	h.Spin()

}

var hzwhertzmw1 = func(ctx context.Context, c *app.RequestContext) {
	slog.Info("=====hzwhertzmw1=====")
}

var hzwhertzmw2 = func(ctx context.Context, c *app.RequestContext) {
	slog.Info("=====hzwhertzmw2 start=====")
	c.Next(ctx)
	slog.Info("=====hzwhertzmw2 end  =====")
}

var hzwhertzmw3 = func(ctx context.Context, c *app.RequestContext) {
	slog.Info("=====hzwhertzmw3 start=====")
	// c.Next(ctx)
	// slog.Info("=====hzwhertzmw3 end  =====")
}

func pongf(msg string) func(ctx context.Context, c *app.RequestContext) {
	return func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": msg})
	}
}
