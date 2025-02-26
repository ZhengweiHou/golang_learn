package sonictest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"testing"

	"github.com/bytedance/sonic"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var dTKSBufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func Test1(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	s1 := slice[2:]
	s2 := slice[:0]

	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)

	bt := []byte{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", bt)

	buff := &bytes.Buffer{}
	buff.Write(bt)
	fmt.Printf("%v\n", buff.Bytes())

	buff.Reset()
	fmt.Printf("%v\n", buff.Bytes())
}

func Test2(t *testing.T) {
	bt := []byte{1, 2, 3, 4, 5, 6}

	buff := &bytes.Buffer{}

	buff.Write(bt)
	sil1 := buff.Bytes()     // sil1 的指针和 buff.buf 一致
	fmt.Printf("%v\n", sil1) // 此时 sil1 的内容确实为1，2，3，4，5，6

	buff.Reset()
	bt2 := []byte{9, 8, 7, 6}
	buff.Write(bt2)      // 此时buff.buf的前四位被替换成了 9，8，7，6
	sil2 := buff.Bytes() // sil2 值为 9，8，7，6  此处的sil2 和上面 sil1 的第一个元素的地址是相同的
	fmt.Printf("%v\n", sil2)
	fmt.Printf("%v\n", sil1)
}

func Test3(t *testing.T) {
	map1 := make(map[string]int, 0)
	for i := 0; i < 10; i++ {
		map1[fmt.Sprintf("k%d", i)] = i
	}

	// sonic 序列化map，key 顺序不能保持一致
	strSonic, _ := sonic.Marshal(map1)
	fmt.Printf("%s\n", strSonic)
	strSonic, _ = sonic.Marshal(map1)
	fmt.Printf("%s\n", strSonic)

	// json 序列化key能保持一致
	strJSON, _ := json.Marshal(map1)
	fmt.Printf("%s\n", strJSON)
	strJSON, _ = json.Marshal(map1)
	fmt.Printf("%s\n", strJSON)

}

func Test4(t *testing.T) {
	map1 := make(map[string]interface{}, 0)
	map1["k1"] = "hello"
	map1["k2"] = "孙悟空"

	strSonic, _ := sonic.Marshal(map1)
	fmt.Printf("%s\n", strSonic)

	reader := transform.NewReader(bytes.NewBuffer(strSonic), simplifiedchinese.GBK.NewEncoder())
	newstr, _ := io.ReadAll(reader)
	fmt.Printf("%s\n", newstr)

}
