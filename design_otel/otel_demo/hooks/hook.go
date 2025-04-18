package hooks

import (
	"fmt"
	"log/slog"

	"github.com/alibaba/opentelemetry-go-auto-instrumentation/pkg/api"
)

func onEnterHelloOTEL3(call api.CallContext, msg string) {
	slog.Info(fmt.Sprintf("otelHook onEnterHelloOTEL3 =%s===", msg))
}
