// package main

// import (
// 	"fmt"

// 	ts "github.com/ShawnRong/tushare-go"
// )

// const apikey = "43e761f1dc45b9a490dcc0e21f138ad0067f9343178049a056d8ddd9"

// func main() {
// 	c := ts.New(apikey)
// 	// 参数
// 	params := make(map[string]string)
// 	params["new_share"] = "L"
// 	// 字段
// 	var fields []string
// 	// 根据api 请求对应的接口
// 	data, _ := c.StockBasic(params, fields)
// 	fmt.Printf("%v\n", data)
// }
