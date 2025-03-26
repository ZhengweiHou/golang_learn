package server

import (
	"github.com/gin-gonic/gin"
	"wire-hzw/internal/handler"
	"wire-hzw/internal/middleware"
	"wire-hzw/pkg/helper/resp"
	"wire-hzw/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	userHandler *handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Nunu!",
		})
	})
	r.GET("/user", userHandler.GetUserById)

	return r
}
