// Code generated by Kitex v0.12.3. DO NOT EDIT.

package hello

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	hzwapi "kitex_demo/kitex_gen/hzwapi"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"echo": kitex.NewMethodInfo(
		echoHandler,
		newHelloEchoArgs,
		newHelloEchoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	helloServiceInfo                = NewServiceInfo()
	helloServiceInfoForClient       = NewServiceInfoForClient()
	helloServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return helloServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return helloServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return helloServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "Hello"
	handlerType := (*hzwapi.Hello)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "hzwapi",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.3",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*hzwapi.HelloEchoArgs)
	realResult := result.(*hzwapi.HelloEchoResult)
	success, err := handler.(hzwapi.Hello).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloEchoArgs() interface{} {
	return hzwapi.NewHelloEchoArgs()
}

func newHelloEchoResult() interface{} {
	return hzwapi.NewHelloEchoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, req *hzwapi.Request) (r *hzwapi.Response, err error) {
	var _args hzwapi.HelloEchoArgs
	_args.Req = req
	var _result hzwapi.HelloEchoResult
	if err = p.c.Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
