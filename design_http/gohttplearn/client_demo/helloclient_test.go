package clientdemo

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"testing"
	"time"
)

func TestHelloClient1(t *testing.T) {
	client := &http.Client{
		Timeout: 3600 * time.Second,
	}
	doHello(client)
	defer client.CloseIdleConnections()
}

// 客户端修改keepalive策略
func TestKeepalive(t *testing.T) {

	// transport := &http.Transport{}
	transport := http.DefaultTransport.(*http.Transport)

	// 自定义RoundTripper
	hzwRt := &custHzwRoundTripper{
		child: transport,
	}

	client := &http.Client{
		Timeout: 3600 * time.Second,
		// Transport: transport,
		Transport: hzwRt,
	}
	defer client.CloseIdleConnections()

	// transport.DisableKeepAlives = true
	// doHello(client)

	transport.DisableKeepAlives = false
	doHello(client)
}

func doHello(client *http.Client) {
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8888/hello", nil)
	q := req.URL.Query()
	for i := 0; i < 3; i++ {
		// 动态添加id请求参数
		q.Set("id", fmt.Sprintf("%d", i))
		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			slog.Error("do request", "err", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			slog.Error("status code", "status", resp.StatusCode)
			return
		}

		body, err := io.ReadAll(resp.Body)
		slog.Info("body", "body", string(body), "err", err)
		time.Sleep(1 * time.Second)
	}

}

type custHzwRoundTripper struct {
	child http.RoundTripper
}

func (h *custHzwRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("cust hzw round tripper")
	return h.child.RoundTrip(req)
}
