package kservice

import (
	"kitex_demo/api/kitex/hello/helloservice"
	"kitex_demo/api/kitex/hzw/hzwcmdservice"
	"kitex_demo/api/kitex/hzw/hzwqueryservice"
	"log"

	server "github.com/cloudwego/kitex/server"
)

func main() {
	svr := server.NewServer()
	if err := hzwcmdservice.RegisterService(svr, new(HzwCmdServiceImpl)); err != nil {
		panic(err)
	}
	if err := hzwqueryservice.RegisterService(svr, new(HzwQueryServiceImpl)); err != nil {
		panic(err)
	}

	if err := helloservice.RegisterService(svr, new(HelloServiceImpl)); err != nil {
		panic(err)
	}

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
