package kservice

import (
	"context"
	"fmt"
	"kitex_demo/api/kitex/helloproto"
	"log/slog"
	"strings"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

// HelloProtoServiceImpl implements the last service interface defined in the IDL.
type HelloProtoServiceImpl struct{}

// Echo implements the HelloProtoServiceImpl interface.
func (s *HelloProtoServiceImpl) Echo(ctx context.Context, req *helloproto.Request) (resp *helloproto.Response, err error) {

	//  String msg = request.getMessage();
	//  System.out.println("get msg:" + msg);
	//  String message = "Hello, we are java!";

	//  if (msg != null && msg.startsWith("err")) {
	// 	 throw new RuntimeException("err");
	//  }

	msg := req.GetMessage()
	slog.Info(fmt.Sprintf("kitex 服务接收到消息:%s", msg))

	// 如果msg 以 err 开头，就抛出异常
	if strings.HasPrefix(msg, "err") {
		// return nil, fmt.Errorf("err")
		return nil, kerrors.NewBizStatusError(555, "服务端估计制造的异常err") // 业务异常
	}

	resp = &helloproto.Response{
		Message: "hello we are golang kitex",
	}
	return
}
