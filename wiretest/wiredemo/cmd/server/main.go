package main

import (
	"context"
	"flag"
	"log/slog"

	"wiredemo/cmd/server/wire"
	"wiredemo/pkg/config"
)

// @title           wire demo API
// @version         1.0.0
// @description     This is a sample server celler server.
/*
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8000
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
*/
func main() {
	var envConf = flag.String("conf", "app.yml", "config path, eg: -conf ./config/app.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	//	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewWire(conf)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	slog.Info("=============starting==========")
	//app.Logger.Info("server start", zap.String("host", fmt.Sprintf("http://%s:%d", conf.GetString("http.host"), conf.GetInt("http.port"))))
	// logger.Info("docs addr", zap.String("addr", fmt.Sprintf("http://%s:%d/swagger/index.html", conf.GetString("http.host"), conf.GetInt("http.port"))))
	//app.Logger.Info("server start", conf.GetString("http.host"), conf.GetInt("http.port"))
	// app.Logger.Info("server start", "host", fmt.Sprintf("http://%s:%d", conf.GetString("http.host"), conf.GetInt("http.port")))
	// app.Logger.Info("docs addr", "addr", fmt.Sprintf("http://%s:%d/swagger/index.html", conf.GetString("http.host"), conf.GetInt("http.port")))
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
