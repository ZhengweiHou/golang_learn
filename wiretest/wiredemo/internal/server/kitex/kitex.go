package kitex

import (
	"fmt"
	"log/slog"
	"net"
	"wiredemo/api/kitex/hzw"
	"wiredemo/api/kitex/hzw/hzwservice"
	"wiredemo/pkg/server/kitex"

	"github.com/cloudwego/kitex/pkg/remote/codec/thrift"
	"github.com/cloudwego/kitex/pkg/remote/connpool"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/spf13/viper"
)

// NewKitexServer 创建程序包装server
func NewKitexServer(
	oriserver server.Server,
	logger *slog.Logger,
	creporter *kitex.HzwKCReporter,
	hzwkctl hzw.HzwService,
) *kitex.Server {

	// 注册hzw kitex 控制器
	//hzwservice.RegisterService(oriserver, hzwkctl)
	hzwservice.RegisterService(oriserver, hzwkctl)

	//oriserver.RegisterService(svcInfo *serviceinfo.ServiceInfo, handler interface{}, opts ...server.RegisterOption)

	kserver := kitex.NewServer(oriserver, logger)

	slog.Info("SerReporter", creporter)
	connpool.SetReporter(creporter) // 注册自定义连接池监控
	return kserver
}

// NewKitexOriginalServer 创建原始的kitex服务实例
func NewKitexOriginalServer(
	logger *slog.Logger,
	conf *viper.Viper,
) server.Server {

	host := conf.GetString("kitex.host")
	port := conf.GetInt("kitex.port")

	kitexHostStr := fmt.Sprintf("%s:%d", host, port)
	logger.Info("kitex", "host", kitexHostStr)

	addr, _ := net.ResolveTCPAddr("tcp", kitexHostStr)
	svr := server.NewServer(
		server.WithServiceAddr(addr),
		// 配合 SkipDecoder https://www.cloudwego.io/zh/docs/kitex/tutorials/code-gen/skip_decoder/
		server.WithPayloadCodec(thrift.NewThriftCodecWithConfig(thrift.FastWrite|thrift.FastRead|thrift.EnableSkipDecoder)),
		//server.WithMuxTransport(),                               // 服务端开启多路复用；Server 开启连接多路复用对 Client 没有限制，可以接受短连接、长连接池、连接多路复用的请求
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),    // 指定基于HTTP2 协议header的信息透传
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler), // 指定基于TTHeader 协议header的信息透传
	)

	return svr
}
