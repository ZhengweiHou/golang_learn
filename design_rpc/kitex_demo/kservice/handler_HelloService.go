package kservice

import (
	"context"
	"fmt"
	"kitex_demo/api/kitex/hello"
	"log/slog"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// Echo implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) Echo(ctx context.Context, req *hello.Request) (resp *hello.Response, err error) {
	slog.Info(fmt.Sprintf("req msg:%s", req.GetMessage()))

	resp = &hello.Response{
		Message: "hello from helloServiceImpl",
	}
	return
}
