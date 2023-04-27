package main

import (
	"flag"
	"fmt"
)

var port = flag.Int("port", 8080, "service port")
var url = flag.String("url", "cron", "schedule URL")
var debug = flag.Bool("debug", false, "debug flag")
var scheduleURLRefreshMs = flag.Int("urlRefresh", 100, "scheduleURL refresh in ms")
var statsHistory = flag.Int("statsHistory", 10, "max stats history")

func main() {
	flag.Parse()

	fmt.Println(*debug, *port)
}
