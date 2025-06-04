package kitexdemo

import (
	"context"
	"fmt"
	"kitex_demo/api/kitex/helloproto"
	"kitex_demo/api/kitex/helloproto/helloprotoservice"
	"kitex_demo/hzwnacosext"
	"kitex_demo/kservice"
	"log"
	"net"
	"strings"
	"testing"
	"time"

	// https://github.com/cloudwego/kitex/issues/1377

	"github.com/cloudwego/kitex/client"
	_ "github.com/cloudwego/kitex/pkg/remote/codec/protobuf/encoding/gzip"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	// kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	// kitexslog "github.com/kitex-contrib/obs-opentelemetry/logging/slog"
)

var hzwendpoint = "hzw-grpc-server" //server name

// ##### CLIENT #####
func TestHelloProtoClientNacos(t *testing.T) {
	naming_client := newNamingClientg()
	helloclient := helloprotoservice.MustNewClient(
		hzwendpoint,
		// client.WithResolver(resolver.NewNacosResolver(naming_client)),
		client.WithResolver(hzwnacosext.NewNacosResolver(naming_client)),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithRPCTimeout(30*time.Second),
	)

	hreq := &helloproto.Request{
		Message: "hello proto golang kitex",
	}
	ai := 0
	ij := 0
	ig := 0
	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Second)
		ai++
		resp, err := helloclient.Echo(context.Background(), hreq)
		if err != nil {
			fmt.Printf("err:%s\n", err.Error())
			continue
		}
		msg := resp.GetMessage()
		if strings.Contains(msg, "java") {
			ij++
		} else if strings.Contains(msg, "golang") {
			ig++
		}
		// fmt.Printf("resp msg:%s\n", resp.GetMessage())
		fmt.Printf("java:%d, golang:%d, all:%d, j+g:%d\n", ij, ig, ai, ij+ig)
	}
	fmt.Printf("java:%d, golang:%d, all:%d, j+g:%d\n", ij, ig, ai, ij+ig)

	// fmt.Printf("java:%d, golang:%d\n", ij, ig)

}

// == kitex server ==
func TestHelloProtoServerNacos(t *testing.T) {
	naming_client := newNamingClientg()
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")

	options := []server.Option{
		server.WithServiceAddr(addr),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: hzwendpoint}),
		// server.WithStatsLevel(stats.LevelDetailed),
		server.WithStatsLevel(stats.LevelDisabled),
	}

	options = append(options, server.WithRegistry(registry.NewNacosRegistry(naming_client)))
	svr := server.NewServer(options...)

	// helloprotoservice.RegisterService(svr, new(kservice.HelloProtoServiceImpl))

	sinfo := helloprotoservice.NewServiceInfo()
	simpl := new(kservice.HelloProtoServiceImpl)
	svr.RegisterService(sinfo, simpl)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func newNamingClientg() naming_client.INamingClient {

	// 创建clientConfig
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId("test"), //当namespace是public时，此处填空字符串。
		// constant.WithTimeoutMs(5000),
		// constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("hzwnacos/log"),
		constant.WithCacheDir("hzwnacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// 创建serverConfig
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			"127.0.0.1",
			8813,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}

	// 创建服务发现客户端
	_, _ = clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		panic(err)
	}

	return namingClient

}
