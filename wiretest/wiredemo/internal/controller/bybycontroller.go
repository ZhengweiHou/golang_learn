package controller

import (
	"context"
	"fmt"
	"wiredemo/internal/service"
	api "wiredemo/pkg/kitex/kitex_gen/api"
)

// BybyImpl implements the last service interface defined in the IDL.
type BybyController struct {
	userService service.IHzwService
}

// Echo implements the BybyImpl interface.
func (s *BybyController) Byby(ctx context.Context, req *api.BybyRequest) (resp *api.BybyResponse, err error) {
	// TODO: Your code here...
	fmt.Println("bbyby")
	return
}

func NewBybyController(userService service.IHzwService) *BybyController {
	return &BybyController{
		userService: userService,
	}
}
