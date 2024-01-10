package base

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	var a *string
	fmt.Print(a)
	if a != nil {
		fmt.Print(a)
	}
}

var h *hBean

func TestPoint2(t *testing.T) {
	h = &hBean{
		name: "hzw",
	}

	htemp := *h             // 从指针获取实际值的拷贝
	htemp.name = "123"      // 拷贝的修改不会影响原值
	fmt.Println(h.name)     // 输出:hzw
	fmt.Println(htemp.name) // 输出:123

}

type hBean struct {
	name string
}

func TestPoint3(t *testing.T) {
	oi := new(hBean) // new 返回的是指针
	oi.name = "h1"

	ni := new(hBean)
	*ni = *oi // 内存复制
	ni.name = "h2"

	fmt.Println(oi.name)
	fmt.Println(ni.name)
}
func TestPoint4(t *testing.T) {

	b1 := &hBean{
		name: "1",
	}

	b2 := *b1
	b1.name = "2"
	fmt.Println(b1.name)
	fmt.Println(b2.name)
}
func TestPoint5(t *testing.T) {
	bs := make([]*hBean, 0)
	for i := 1; i < 10; i++ {
		bs = append(bs, &hBean{
			name: fmt.Sprintf("hbean_%d", i),
		})
	}

	var wg sync.WaitGroup

	for _, _hb := range bs {
		wg.Add(1)
		go func(hb *hBean) { // 该方式不会被range影响
			oname := hb.name
			time.Sleep(time.Second * 2)
			fmt.Printf("oname:%s, name:%s\n", oname, hb.name)
			wg.Done()
		}(_hb)
	}

	time.Sleep(time.Second)
	a := bs[3]
	a.name = "xiugaile" // 若通过具体指针修改，会影响上面对应的元素

	wg.Wait()

	fmt.Println("========")
	for _, _hb := range bs {
		wg.Add(1)
		go func() {
			oname := _hb.name // 这个用法会被range影响
			time.Sleep(time.Second * 1)
			fmt.Printf("oname:%s, name:%s\n", oname, _hb.name)
			wg.Done()
		}()
	}
	wg.Wait()
}
