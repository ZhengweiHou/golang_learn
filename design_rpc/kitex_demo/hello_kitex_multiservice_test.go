package kitexdemo

import (
	"context"
	"fmt"
	"kitex_demo/api/kitex/hello"
	"kitex_demo/api/kitex/hello/helloservice"
	"kitex_demo/api/kitex/helloproto"
	"kitex_demo/api/kitex/helloproto/helloprotoservice"
	"kitex_demo/api/kitex/hzw"
	"kitex_demo/api/kitex/hzw/hzwcmdservice"
	"kitex_demo/api/kitex/hzw/hzwqueryservice"
	"kitex_demo/kservice"
	"log"
	"log/slog"
	"net"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/server"
)

func TestMutilServiceServer(t *testing.T) {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	svr := server.NewServer(
		server.WithServiceAddr(addr),
	)

	err := helloservice.RegisterService(svr, new(kservice.HelloServiceImpl))
	if err != nil {
		slog.Error(err.Error())
	}
	err = hzwcmdservice.RegisterService(svr, new(kservice.HzwCmdServiceImpl))
	if err != nil {
		slog.Error(err.Error())
	}
	err = hzwqueryservice.RegisterService(svr, new(kservice.HzwQueryServiceImpl))
	if err != nil {
		slog.Error(err.Error())
	}
	//	helloprotoservice.RegisterService(svr, new(kservice.HelloProtoServiceImpl)) // 方式1
	err = svr.RegisterService(helloprotoservice.NewServiceInfo(), new(kservice.HelloProtoServiceImpl)) // 方式2
	if err != nil {
		slog.Error(err.Error())
	}

	sinfos := svr.GetServiceInfos()
	// 格式化打印服务信息
	for sname, sinfo := range sinfos {
		slog.Info("Kitex服务信息",
			"服务名称", sname,
			"方法数量", len(sinfo.Methods),
			"方法列表", sinfo.Methods,
			"额外信息", sinfo.Extra,
		)
	}

	slog.Info("Kitex服务启动",
		"总服务数", len(sinfos),
	)

	err = svr.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func TestMutilServcieClient(t *testing.T) {
	// 客户端调用没有变化
	c, _ := helloservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	req := &hello.Request{Message: "hello request"}
	resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resp msg:%s\n", resp.GetMessage())

	c2, _ := hzwcmdservice.NewClient("hzw", client.WithHostPorts("0.0.0.0:8888"))
	req2 := &hzw.Request{Message: "hello hzw"}
	resp2, err := c2.Echo2(context.Background(), req2)
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Printf("resp msg:%s\n", resp2.GetMessage())

	c3, _ := hzwqueryservice.NewClient("hzw", client.WithHostPorts("0.0.0.0:8888"))
	req3 := &hzw.Request{Message: "hello hzw"}
	resp3, err := c3.Echo1(context.Background(), req3)
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Printf("resp msg:%s\n", resp3.GetMessage())

	c4, _ := helloprotoservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	req4 := &helloproto.Request{Message: "hello proto"}
	resp4, err := c4.Echo(context.Background(), req4)
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Printf("resp msg:%s\n", resp4.GetMessage())
}
