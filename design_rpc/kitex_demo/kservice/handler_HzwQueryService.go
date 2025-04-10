package kservice

import (
	"context"
	"fmt"
	"kitex_demo/api/kitex/hzw"
	"log/slog"
)

// HzwQueryServiceImpl implements the last service interface defined in the IDL.
type HzwQueryServiceImpl struct{}

// Echo1 implements the HzwQueryServiceImpl interface.
func (s *HzwQueryServiceImpl) Echo1(ctx context.Context, req *hzw.Request) (resp *hzw.Response, err error) {
	slog.Info(fmt.Sprintf("req msg:%s", req.GetMessage()))

	resp = &hzw.Response{
		Message: "hello from hzwquery Impl",
	}
	return
}
