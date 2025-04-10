package kservice

import (
	"context"
	"fmt"
	"kitex_demo/api/kitex/helloproto"
	"log/slog"
)

// HelloProtoServiceImpl implements the last service interface defined in the IDL.
type HelloProtoServiceImpl struct{}

// Echo implements the HelloProtoServiceImpl interface.
func (s *HelloProtoServiceImpl) Echo(ctx context.Context, req *helloproto.Request) (resp *helloproto.Response, err error) {
	slog.Info(fmt.Sprintf("kitex server get req msg:%s", req.GetMessage()))

	resp = &helloproto.Response{
		Message: "hello we are golang kitex",
	}
	return
}
