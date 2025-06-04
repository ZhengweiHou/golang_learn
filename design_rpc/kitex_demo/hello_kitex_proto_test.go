package kitexdemo

import (
	"context"
	"fmt"
	"kitex_demo/api/kitex/helloproto"
	"kitex_demo/api/kitex/helloproto/helloprotoservice"
	"kitex_demo/kservice"
	"log"
	"net"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"

	// https://github.com/cloudwego/kitex/issues/1377
	"github.com/cloudwego/kitex/pkg/klog"
	_ "github.com/cloudwego/kitex/pkg/remote/codec/protobuf/encoding/gzip"
	"github.com/cloudwego/kitex/pkg/transmeta"

	// kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	// kitexslog "github.com/kitex-contrib/obs-opentelemetry/logging/slog"
	kitexslog "github.com/kitex-contrib/obs-opentelemetry/logging/slog"
)

func TestHK2JavaClient(t *testing.T) {
	// helloProtoCall(8082)
	// helloProtoCall(8099)
	helloProtoCall(19883)
}

func TestHK2KitexClient(t *testing.T) {
	// helloProtoCall(8888)
	helloProtoCall(8001)
}

func helloProtoCall(port int) {
	c, err := helloprotoservice.NewClient(
		"hello",
		client.WithHostPorts(fmt.Sprintf("0.0.0.0:%d", port)),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	if err != nil {
		log.Fatal(err)
	}
	hreq := &helloproto.Request{
		Message: "hello proto golang kitex",
	}
	resp, err := c.Echo(context.Background(), hreq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resp msg:%s\n", resp.GetMessage())
}

// == kitex server ==
func TestHelloProtoServer(t *testing.T) {
	// klog.SetLogger(kitexzap.NewLogger())
	klog.SetLogger(kitexslog.NewLogger())

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	svr := server.NewServer(
		server.WithServiceAddr(addr),
	)

	// svr := helloprotoservice.NewServer(new(kservice.HelloProtoServiceImpl))
	svr.RegisterService(helloprotoservice.NewServiceInfo(), new(kservice.HelloProtoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
