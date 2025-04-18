package hellodo

import (
	"fmt"
	"log/slog"
)

type Hello struct {
}

func (h *Hello) HelloOTEL3(msg string) string {
	slog.Info("hello otel3====== ", "msg", msg)
	return fmt.Sprintf("resp %s", msg)
}
