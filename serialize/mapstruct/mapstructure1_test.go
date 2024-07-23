package mapstruct

import (
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
)

func Test1(t *testing.T) {
	type SASL struct {
		Enable bool
	}
	type User struct {
		ID   string `json:"id"`
		Name string `json:"name,omitempty"`
		Sasl SASL   `mapstructure:"sasl"`
	}
	user := &User{
		ID:   "1",
		Name: "hzw",
		Sasl: SASL{
			Enable: true,
		},
	}

	map1 := make(map[string])

	dec := &mapstructure.DecoderConfig{
		Result: &map1,
	}

	decoder, err := mapstructure.NewDecoder(dec)
	if err != nil {
		t.Fatal(err)
	}
	decoder.Decode(user)
	fmt.Printf("%v\n", map1)
}
