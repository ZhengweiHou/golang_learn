package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"
)

type YmlBean1 struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username,omitempty" json:"name,omitempty"`
	Password string `yaml:"password"`
}

func Test_yaml1(t *testing.T) {

	db := &YmlBean1{
		Host:     "123",
		Password: "abc",
	}

	d, _ := yaml.Marshal(db)
	fmt.Printf("%s\n", string(d))

	d2, _ := json.Marshal(db)
	fmt.Printf("%s\n", string(d2))
}
