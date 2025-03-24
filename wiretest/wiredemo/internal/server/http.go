package server

import (
	apiV1 "wiredemo/api/http/v1"
	"wiredemo/docs"
	"wiredemo/internal/controller"
	"wiredemo/pkg/log"
	"wiredemo/pkg/server/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	hzwController *controller.HzwController,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = apiV1.VBASE
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "hello wire demo",
		})
	})

	vbase := apiV1.VBASE
	vapi := s.Group(vbase)
	{
		// 不需要鉴权的router
		noAuthRouter := vapi.Group("/")
		{
			noAuthRouter.GET("/hello", func(ctx *gin.Context) {
				logger.WithContext(ctx).Info("hello")
				apiV1.HandleSuccess(ctx, map[string]interface{}{
					":)": "hello wire demo",
				})
			})
		}

		// 不需要严格鉴权的router
		noStrictAuthRouter := vapi.Group("/").Use(
		//middleware.NoStrictAuth(jwt, logger) // TODO添加鉴权等中间件操作
		)
		{
			noStrictAuthRouter.GET("/hzw", hzwController.QueryById)
		}

		// 需要严格鉴权的router
		strictAuthRouter := vapi.Group("/").Use(
		//middleware.StrictAuth(jwt, logger) // TODO添加鉴权等中间件操作
		)
		{
			strictAuthRouter.PUT("/hzw", hzwController.SaveHzw)
			strictAuthRouter.PUT("/hzwtxtest", hzwController.SaveHzwTxTest)
			strictAuthRouter.PUT("/hzwwithtx", hzwController.SaveHzwWithTx)
		}
	}

	return s
}

//func NewServerHTTP(
//	hzwController *controller.HzwController,
//) *gin.Engine {
//	gin.SetMode(gin.ReleaseMode)
//	r := gin.Default()
//	r.GET("/", func(ctx *gin.Context) {
//		resp.HandleSuccess(ctx, map[string]interface{}{
//			"say": "Hello!!! 你好!!!",
//		})
//	})
//	r.GET("/hzw", hzwController.QueryById)
//	r.POST("/savehzw", hzwController.SaveHzw)
//
//	// swagger doc
//	//docs.SwaggerInfo.BasePath = "/v1"
//	docs.SwaggerInfo.BasePath = "/"
//	r.GET("/swagger/*any", ginSwagger.WrapHandler(
//		swaggerfiles.Handler,
//		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
//		ginSwagger.DefaultModelsExpandDepth(-1),
//		ginSwagger.PersistAuthorization(true),
//	))
//
//	return r
//}
