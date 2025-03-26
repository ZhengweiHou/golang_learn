package hello

import (
	"context"
	"fmt"
	hzwapi "kitex_demo/kitex_gen/hzwapi"
)

// HelloImpl implements the last service interface defined in the IDL.
type HelloImpl struct{}

// Echo implements the HelloImpl interface.
func (s *HelloImpl) Echo(ctx context.Context, req *hzwapi.Request) (resp *hzwapi.Response, err error) {
	// TODO: Your code here...
	fmt.Printf("recive req msg:%s\n", req.Message)
	resp = hzwapi.NewResponse()
	resp.Message = "hello kitex hzw"
	return
}
