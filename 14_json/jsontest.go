package main

import (
	"encoding/json"
	"fmt"
)

// type User1 struct {
// 	UserId   string
// 	UserName string
// 	age      int
// 	sex      string
// }

func main() {
	// == T1 ==
	type User1 struct {
		UserId   string
		UserName string
		age      int
		sex      string
	}
	u := User1{
		UserId: "1",
		sex:    "男",
	}
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("大写打印,空值取默认; 小写忽略,默认使用变量名\n%s\n", string(data))

	// == T2 ==
	type User2 struct {
		UserId   string `json:"id"`
		UserName string `json:"name"`
		age      int    `json:"age"`
		sex      string `json:"sex"`
	}
	u2 := User2{
		UserId: "1",
		sex:    "男",
	}
	data2, _ := json.Marshal(u2)
	fmt.Printf("规则同上，使用json标记名称\n%s\n", string(data2))

	// == T3 ==
	type User3 struct {
		UserId   string `json:"id"`
		UserName string `json:"name,omitempty"`
		age      int    `json:"age,omitempty"`
		Sex      string `json:"sex,omitempty"`
	}
	u3 := User3{
		UserId: "1",
		Sex:    "男",
		age:    12,
	}
	data3, _ := json.Marshal(u3)
	fmt.Printf("忽略小写，omitempty标记时，忽略空值字段\n%s\n", string(data3))

	// == T4 ==
	type User4 struct {
		UserId   string `json:"id"`
		UserName string `json:"name,omitempty"`
		Age      int    `json:"-"`
		Sex      string `json:"sex,omitempty"`
	}
	u4 := User4{
		UserId:   "1",
		UserName: "张三",
		Age:      12,
	}
	data4, _ := json.Marshal(u4)
	fmt.Printf("- 标记的字段忽略\n%s\n", string(data4))

}
