package kitexdemo

import (
	"context"
	"fmt"
	"kitex_demo/api/kitex/hello"
	"kitex_demo/api/kitex/hello/helloservice"
	"kitex_demo/kservice"
	"log"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func TestHelloServer(t *testing.T) {
	// svr := hello.NewServer(new(rpc.HelloImpl))
	svr := helloservice.NewServer(new(kservice.HelloServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func TestHelloClient(t *testing.T) {
	c, err := helloservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	hreq := hello.NewRequest()
	hreq.Message = "hello req"

	resp, err := c.Echo(context.Background(), hreq, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("resp msg:%s\n", resp.GetMessage())

}
