package kitex

import "context"

// HelloReq是请求结构体
type HelloReq struct {
	Name string
}

// HelloResp是响应结构体
type HelloResp struct {
	Message string
}

// HelloService是服务接口
type HelloService interface {
	Hello(ctx context.Context, req *HelloReq) (*HelloResp, error)
}

// HelloServiceImpl是服务接口的实现
type HelloServiceImpl struct{}

// Hello方法实现了HelloService接口
func (s *HelloServiceImpl) Hello(ctx context.Context, req *HelloReq) (*HelloResp, error) {
	return &HelloResp{Message: "Hello, " + req.Name}, nil
}
