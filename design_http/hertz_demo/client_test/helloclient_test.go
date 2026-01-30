package clienttest

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func TestHelloHertzClient(t *testing.T) {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	req := &protocol.Request{}
	req.SetMethod(consts.MethodGet)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI("http://localhost:8888/ping")

	res := &protocol.Response{}
	err = c.Do(context.Background(), req, res)
	if err != nil {
		return
	}
	fmt.Printf("%v", string(res.Body()))
}

func TestHertzClient(t *testing.T) {
	c, err := client.NewClient(
		client.WithMaxConnsPerHost(10),
		// client.WithMaxConnDuration()
		client.WithConnStateObserve(func(s config.HostClientState) {
			pools := s.ConnPoolState()
			fmt.Printf("%v", pools)
		}),
	)
	if err != nil {
		return
	}

	// c.Use(sd.Discovery())

	req := &protocol.Request{}
	req.SetMethod(consts.MethodGet)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI("http://localhost:8888/ping")

	res := &protocol.Response{}
	err = c.Do(context.Background(), req, res)
	if err != nil {
		return
	}
	fmt.Printf("%v", string(res.Body()))

}
