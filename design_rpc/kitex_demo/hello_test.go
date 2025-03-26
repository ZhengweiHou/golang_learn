package main

import (
	"context"
	"fmt"
	"kitex_demo/kitex_gen/hzwapi"
	helloapi "kitex_demo/kitex_gen/hzwapi/hello"
	"kitex_demo/rpc/hello"
	"log"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func TestHelloServer(t *testing.T) {
	svr := helloapi.NewServer(new(hello.HelloImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func TestHelloClient(t *testing.T) {
	c, err := helloapi.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	hreq := hzwapi.NewRequest()
	hreq.Message = "hello req"

	resp, err := c.Echo(context.Background(), hreq, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("resp msg:%s\n", resp.GetMessage())

}
