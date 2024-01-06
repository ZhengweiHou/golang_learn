package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_json2_1(t *testing.T) {

	type user struct {
		Name string
		Time time.Time
		Date time.Time
	}

	u := user{
		Name: "hzw",
		Time: time.Now(),
		Date: time.Now(),
	}
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n", string(data))

}

func Test_json2_2(t *testing.T) {
	map1 := make(map[string]interface{}, 0)
	map1["name"] = "hzw"
	map1["time"] = time.Now()
	map1["date"] = time.Now()
	fmt.Println("=原始map")
	for k, v := range map1 {
		fmt.Printf("k:%s, t:%T, v:%v\n", k, v, v)
	}
	d, err := json.Marshal(map1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("=原始map json")
	fmt.Printf("%s\n", string(d))

	map2 := make(map[string]interface{}, 0)
	map2["name"] = "hzw2"
	map2["time"] = time.Now()
	map2["date"] = time.Now()

	fmt.Println("=反序列化前模板map")
	for k, v := range map2 {
		fmt.Printf("k:%s, t:%T, v:%v\n", k, v, v)
	}

	err = json.Unmarshal(d, &map2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("=反序列化后模板map")
	for k, v := range map2 {
		fmt.Printf("k:%s, t:%T, v:%v\n", k, v, v)
	}
}
