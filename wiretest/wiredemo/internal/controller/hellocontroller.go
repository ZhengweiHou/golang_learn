package controller

import (
	"context"
	"fmt"
	"strconv"
	"wiredemo/internal/service"
	api "wiredemo/pkg/kitex/kitex_gen/api"
)

// HelloController implements the last service interface defined in the IDL.
type HelloController struct {
	userService service.IHzwService
}

// Echo implements the HelloImpl interface.
func (s *HelloController) Echo(ctx context.Context, req *api.HelloRequest) (resp *api.HelloResponse, err error) {
	fmt.Println("echo begin")
	id, _ := strconv.Atoi(req.Id)
	hzw, err := s.userService.GetHzw(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("echo end %s\n", hzw.Name)
	resp = &api.HelloResponse{
		Name: hzw.Name,
	}
	return
}

func NewHelloController(userService service.IHzwService) *HelloController {
	return &HelloController{
		userService: userService,
	}
}
