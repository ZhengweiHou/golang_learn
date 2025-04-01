package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
	"wiredemo/api/kitex/hzw"
	"wiredemo/api/kitex/hzw/hzwservice"
	"wiredemo/pkg/server/kitex"

	"github.com/cloudwego/kitex/client"
	connpool2 "github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/remote/connpool"
	"github.com/cloudwego/kitex/transport"
)

func TestKitexTest(t *testing.T) {
	ctx := context.Background()
	c, _ := hzwservice.NewClient(
		"hzw",
		client.WithHostPorts("0.0.0.0:8001"),
		//client.WithMuxConnection(2), // 客户端开启多路复用调用方式
		client.WithTransportProtocol(transport.TTHeader), // Thrift则多路复用只支持TTHeader传输协议，gRPC默认连接是多路复用
		//		client.WithTransportProtocol(transport.GRPC), // Thrift则多路复用只支持TTHeader传输协议，gRPC默认连接是多路复用
		client.WithLongConnection(connpool2.IdleConfig{10, 100, 100, time.Minute}),
		client.WithConnReporterEnabled(),
	)

	// 设置自定义连接池监控
	hzwCReporter := kitex.NewHzwKCReporter()
	connpool.SetReporter(hzwCReporter)

	hzwDto := &hzw.HzwDto{}
	hzwDto.SetName("kitextest")
	hzwDto.SetAge(22)
	rhzw, err := c.CreateHzw(ctx, hzwDto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("rhzw:%s", rhzw.String())
}
