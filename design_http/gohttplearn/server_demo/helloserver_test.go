package serverdemo

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"testing"
	"time"
)

func TestHelloServer1(t *testing.T) {
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", 8888)
	slog.Info("hello server start", "addr", addr)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		par_id := r.URL.Query().Get("id")

		t := time.Now().Format("2006-01-02 15:04:05.000")
		sayreq := fmt.Sprintf("time:%v 接收到请求 id=%s", t, par_id)

		slog.Info(sayreq)
		// slog.Info("hello world", "addr", addr, "time", time.Now().Format("2006-01-02 15:04:05.000"))
		fmt.Fprintf(w, fmt.Sprintf("hello: %s time: %v", par_id, t))
	})

	httpSrv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		// log.Fatalf("listen: %s\n", err)
		// slog.Error("hello server", "err", err)
		slog.Info("hello server", "err", err)
	}
}
