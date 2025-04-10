package kservice

import (
	"context"
	"fmt"
	"kitex_demo/api/kitex/hzw"
	"log/slog"
)

// HzwCmdServiceImpl implements the last service interface defined in the IDL.
type HzwCmdServiceImpl struct{}

// Echo2 implements the HzwCmdServiceImpl interface.
func (s *HzwCmdServiceImpl) Echo2(ctx context.Context, req *hzw.Request) (resp *hzw.Response, err error) {
	slog.Info(fmt.Sprintf("req msg:%s", req.GetMessage()))

	resp = &hzw.Response{
		Message: "hello from hzwcmd Impl",
	}
	return
}
