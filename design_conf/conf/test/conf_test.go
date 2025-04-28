package test

import (
	"design_conf/conf"
	"fmt"
	"testing"
)

func TestConf(t *testing.T) {
	standEnv := conf.NewStandardEnvironment()

	v := standEnv.GetProperty("home")
	fmt.Printf("v: %v\n", v)

}
