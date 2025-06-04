package vipertest

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestVipertest1(t *testing.T) {
	path := "./config.yml"
	cnf := viper.New()
	cnf.SetConfigFile(path)
	err := cnf.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("======")
	fmt.Printf(" %v\n", cnf.AllKeys())
}
