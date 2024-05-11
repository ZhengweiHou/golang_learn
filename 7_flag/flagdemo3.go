package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

// go run flagdemo3.go -a=a -b=b
func main() {

	var a1 string
	flag.StringVar(&a1, "a1", "", "")
	flag.Parse()
	fmt.Printf("a1:%s\n", a1)

	args := flag.Args()

	aj, _ := json.MarshalIndent(args, "", " ")
	fmt.Printf("%s\n", aj)

}
