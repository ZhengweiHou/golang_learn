package kitexdemo

import (
	"context"
	"fmt"
	"kitex_demo/api/kitex/helloproto"
	"kitex_demo/api/kitex/helloproto/helloprotoservice"
	"kitex_demo/kservice"
	"log"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"

	// https://github.com/cloudwego/kitex/issues/1377
	_ "github.com/cloudwego/kitex/pkg/remote/codec/protobuf/encoding/gzip"
	"github.com/cloudwego/kitex/pkg/transmeta"
)

func TestHK2JavaClient(t *testing.T) {
	helloProtoCall(8082)
}

func TestHK2KitexClient(t *testing.T) {
	helloProtoCall(8888)
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
	svr := helloprotoservice.NewServer(new(kservice.HelloProtoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
