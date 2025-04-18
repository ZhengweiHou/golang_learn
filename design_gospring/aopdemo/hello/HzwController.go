package hello

type HController struct {
	svc IHService
}

func NewHController(svc IHService) *HController {
	proxy := NewGenericProxy(svc, func(method string, args []interface{}, result []interface{}) {
		// 可以在这里添加自定义AOP逻辑
	})
	return &HController{
		svc: proxy.(IHService),
	}
}

func (h *HController) HelloCtl(msg string) string {
	return h.svc.HelloSvs(msg)
}
