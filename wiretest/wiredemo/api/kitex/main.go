package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"wiredemo/pkg/kitex/kitex_gen/api"
	"wiredemo/pkg/kitex/kitex_gen/api/hello"
)

func main() {
	cli, err := hello.NewClient("wiredemo-kitex-server", client.WithHostPorts("0.0.0.0:8999"))
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := cli.Echo(context.Background(), &api.HelloRequest{Name: "hello", Id: "1"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
