package designcache

import (
	"encoding/base64"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

var callCount uint64
var sharedCount uint64

var sg = singleflight.Group{}

func TestMain(t *testing.T) {
	// 1000个协程同时请求
	for i := 0; i < 1000; i++ {
		go sfget("123")
	}
	time.Sleep(time.Millisecond * 500)

	fmt.Printf("Total calls: %d, shared: %d\n", atomic.LoadUint64(&callCount), atomic.LoadUint64(&sharedCount))

	time.Sleep(time.Millisecond * 500)

	for i := 0; i < 1000; i++ {
		go sfget("123")
	}
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("Total calls: %d, shared: %d\n", atomic.LoadUint64(&callCount), atomic.LoadUint64(&sharedCount))
}

func sfget(k string) string {
	a, _, shared := sg.Do(k, func() (interface{}, error) {
		// fmt.Printf("do %s\n", k)
		atomic.AddUint64(&callCount, 1)

		str := fmt.Sprintf("hello %s", k)
		return base64.StdEncoding.EncodeToString([]byte(str)), nil
	})
	if shared {
		atomic.AddUint64(&sharedCount, 1)
	}

	return a.(string)

}
