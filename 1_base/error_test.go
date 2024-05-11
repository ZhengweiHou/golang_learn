package base

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr1(t *testing.T) {
	err := ErrH
	err2 := fmt.Errorf("warp %w", err)

	fmt.Printf("is Herror:%v\n", errors.Is(err2, ErrH))
	fmt.Printf("is Herror:%v\n", errors.Is(err, ErrH))
}

var ErrH = errors.New("herror")

func TestErr2(t *testing.T) {
	var err error

	err1 := fmt.Errorf("%s", "error1")
	err = fmt.Errorf("%w err:%w", err, err1)

	err2 := fmt.Errorf("%s", "error2")
	err = fmt.Errorf("%w err:%w", err, err2)

	fmt.Printf("%v\n", err)
	// %!w(<nil>) error1 error2

}
