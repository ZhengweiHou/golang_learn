package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// type User1 struct {
// 	UserId   string
// 	UserName string
// 	age      int
// 	sex      string
// }

func Test1(t *testing.T) {
	type User struct {
		ID    string `json:"id"`
		Name  string `json:"name,omitempty"`
		age   int    `json:"age,omitempty"`   // 首字母小写属性json序列化会被忽略
		Msg   string `json:"sex,omitempty"`   // omitempty 忽略空值
		Bool1 bool   `json:"bool1,omitempty"` // omitempty 在bool为false时为空
		Msg2  string `json:""`                //
	}
	user := &User{
		ID:    "1",
		Name:  "hzw",
		age:   12,   // 首字母小写属性json序列化会被忽略
		Msg:   "",   // omitempty 忽略空值
		Bool1: true, // omitempty 在bool为false时为空
		Msg2:  "hhhhhhh",
	}

	// === 反射获取标签信息 ===
	// 反射获取类型
	utype := reflect.TypeOf(user).Elem()
	// 通过类型获取字段定义
	unamefield, _ := utype.FieldByName("Name")
	// 通过字段定义获取指定名称的tag描述
	tag := unamefield.Tag.Get("json")
	fmt.Printf("通过反射获取到name的json标签:%s\n", tag)

	// === json 序列化 ===
	data3, _ := json.Marshal(user)
	fmt.Printf("%s\n", string(data3))

}

func Test2(t *testing.T) {

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
		UserName string `json:"name,omitempty"` // omitempty 忽略空值
		age      int    `json:"age,omitempty"`
		Sex      string `json:"sex,omitempty"`
		Bool1    bool   `json:"bool1,omitempty"` // omitempty 在bool为false时为空
	}
	u3 := User3{
		UserId: "1",
		Sex:    "男",
		age:    12,
		Bool1:  true,
	}

	utype := reflect.TypeOf(&u3).Elem()
	unamefield, _ := utype.FieldByName("UserName")
	tag := unamefield.Tag.Get("json")
	fmt.Printf("u3.username jsontag:%s", tag)

	data3, _ := json.Marshal(u3)
	fmt.Printf("T3 忽略小写,omitempty标记时忽略空值字段\n%s\n", string(data3))

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
