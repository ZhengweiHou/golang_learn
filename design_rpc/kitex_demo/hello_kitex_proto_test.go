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
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"

	// kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	// kitexslog "github.com/kitex-contrib/obs-opentelemetry/logging/slog"
	kitexslog "github.com/kitex-contrib/obs-opentelemetry/logging/slog"
)

func TestHK2JavaClient(t *testing.T) {
	// helloProtoCall(8082)
	// helloProtoCall(2222)
	// helloProtoCall(8088)
	// helloProtoCall(19883)
	helloProtoCall(7777, false)
	helloProtoCall(8888, false)
}
func Test1(t *testing.T) {
	helloProtoCall(8082, false)
}
func Test2(t *testing.T) {
	helloProtoCall(8886, false)
}

func TestHK2KitexClient(t *testing.T) {
	// helloProtoCall(1111) // gw1
	// helloProtoCall(2222) // gw2
	// helloProtoCall(7777, true) // java
	helloProtoCall(8888, true) // golang
}

func TestClient(t *testing.T) {

}
func TestHelloProtoCall(t *testing.T) {
	c, _ := helloprotoservice.NewClient(
		"hello",
		client.WithHostPorts(fmt.Sprintf("127.0.0.1:%d", 8888)),
		client.WithTransportProtocol(transport.GRPC),
	)
	hreq := &helloproto.Request{
		Message: "hello proto golang kitex",
	}
	resp, _ := c.Echo(context.Background(), hreq)
	fmt.Printf("resp msg:%s\n", resp.GetMessage())
}

func helloProtoCall(port int, e bool) {
	c, err := helloprotoservice.NewClient(
		"hello",
		client.WithHostPorts(fmt.Sprintf("0.0.0.0:%d", port)),
		client.WithTransportProtocol(transport.GRPC),
		// client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"key1", "value1",
	)

	for i := 0; i < 5; i++ {
		hreq := &helloproto.Request{
			Message: fmt.Sprintf("%s %d", "hello proto golang kitex", i),
		}

		if e {
			hreq.Message = "err" + hreq.Message
		}

		resp, err := c.Echo(ctx, hreq)
		if err != nil {
			fmt.Printf("调用返回异常:%T， %s\n", err, err.Error())
			log.Fatal(err)
		}
		fmt.Printf("resp msg:%s\n", resp.GetMessage())
	}

}

// == kitex server ==
func TestHelloProtoServer(t *testing.T) {
	klog.SetLogger(kitexslog.NewLogger())
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	svr := server.NewServer(
		server.WithServiceAddr(addr),
	)
	svr.RegisterService(helloprotoservice.NewServiceInfo(), new(kservice.HelloProtoServiceImpl))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
