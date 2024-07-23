package runtimetest

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGcMap1(t *testing.T) {

	var rmap = make(map[string]interface{})
	rmap["k1"] = "v1"

	rslice := mapPush(rmap)

	rmap["k1"] = "v2"
	rb, _ := json.Marshal(rslice)

	fmt.Printf("%s\n", rb)
}

func mapPush(rmap map[string]interface{}) []interface{} { // 此处map是指针传递
	rlen := len(rmap)

	var rslice = make([]interface{}, rlen) // 数组
	index := 0
	for _, v := range rmap {
		rslice[index] = v
		index++
	}
	return rslice
}
