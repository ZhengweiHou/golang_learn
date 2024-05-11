package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

func main() {
	// 读取配置文件
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("无法读取配置文件：%v", err)
	}

	// 解析配置文件
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("无法解析配置文件：%v", err)
	}

	// 打印配置信息
	fmt.Printf("数据库主机：%s\n", config.Database.Host)
	fmt.Printf("数据库端口：%d\n", config.Database.Port)
	fmt.Printf("数据库用户名：%s\n", config.Database.Username)
	fmt.Printf("数据库密码：%s\n", config.Database.Password)
}
