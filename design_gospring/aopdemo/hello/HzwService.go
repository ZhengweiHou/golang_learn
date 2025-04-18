package hello

import "log/slog"

type IHService interface {
	HelloSvs(msg string) string
}

type HService struct {
}

func NewHService() IHService {
	return &HService{}
}

func (h *HService) HelloSvs(msg string) string {
	slog.Info("Hello", "msg", msg)
	return "HService hello"
}
