package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/axgle/mahonia"
	utils "github.com/bigwhite/gocmpp/utils"
	"github.com/tidwall/gjson"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var numArr = []int{210, 202, 223, 220, 215, 220, 238, 238, 228, 232, 239, 236, 239, 217}

func main() {
	// fmt.Println("Hello, World!")
	// sort.Sort(sort.IntSlice(numArr))

	// for _, v := range numArr {
	// 	fmt.Println(v)
	// }

	// name := "Señor"
	// printUnicode(name)

	//test_for_Range()
	//test_for_gbk_utf8()
	// name := []byte{0xD6, 0xD0, 0xB9, 0xFA}
	// printBytes4(name)
	// a := test_gbk_encoding2()

	// fmt.Println("a is ", a)
	//converts a  string from UTF-8 to gbk encoding.
	//fi, err := os.Open("utf.txt")

	fmt.Printf(test_gbk_encoding())
	fmt.Printf(test_gbk_encoding2())

}

func getMedian() float64 {
	length := len(numArr)
	if length == 0 {
		return -1
	}
	sort.Ints(numArr)
	if length%2 == 0 {
		return float64(numArr[length/2-1]+numArr[length/2]) / 2
	} else {
		return float64(numArr[length/2])
	}
}

func test_for_string() {
	var s = "中"
	fmt.Printf("%s => UTF8编码: ", s)
	for _, v := range []byte(s) {
		fmt.Printf("%X", v)
	}
	fmt.Printf("\n")
	fmt.Printf("%s => Unicode编码: ", s)
	for _, v := range s {
		fmt.Printf("%X", v)
	}
	fmt.Printf("\n")
}

func test_for_gbk_utf8() {
	var stringLiteral = "中国人"
	var stringUsingRuneLiteral = "\u4E2D\u56FD\u4EBA"

	if stringLiteral != stringUsingRuneLiteral {
		fmt.Println("stringLiteral is not equal to stringUsingRuneLiteral")
		return
	}

	for i := 0; i < len(stringLiteral); i++ {
		fmt.Printf("literal is %x  ", stringLiteral[i])
	}

	fmt.Println("stringLiteral is equal to stringUsingRuneLiteral")

	for i, v := range stringLiteral {
		fmt.Printf("中文字符: %s <=> Unicode码点(rune): %X <=> UTF8编码(内存值): ", string(v), v)
		s := stringLiteral[i : i+3]
		for _, v := range []byte(s) {
			fmt.Printf("0x%X ", v)
		}

		s1, _ := utils.Utf8ToGB18030(s)
		fmt.Printf("<=> GB18030编码(内存值): ")
		for _, v := range []byte(s1) {
			fmt.Printf("0x%X ", v)
		}
		fmt.Printf("\n")
	}
}

func printUnicode(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%x ", runes[i])
	}
}

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}
func printBytes2(s string) {
	bytes := []byte(s)

	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%x ", bytes[i])
	}
}

func printBytes3(b []byte) {
	for i := 0; i < len(b); i++ {
		fmt.Printf("%x ", b[i])
	}
}

func printBytes4(b []byte) {
	s := string(b)
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%s ", s[i])
	// }
	fmt.Println(s)
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

// }
func test_gbk_encoding() string {
	enc := mahonia.NewDecoder("gbk")
	buf, err := ioutil.ReadFile("gbk.txt")
	if err != nil {
		fmt.Println("file read error!")
	}
	bufc := enc.ConvertString(string(buf))
	return bufc
}

func test_gbk_encoding2() string {
	buf, err := ioutil.ReadFile("gbk.txt")
	if err != nil {
		fmt.Println("file read error!")
	}
	bufc, _ := simplifiedchinese.GBK.NewDecoder().Bytes(buf)
	return string(bufc)
}

func test_gbk_noencoding() string {
	buf, err := ioutil.ReadFile("gbk.txt")
	if err != nil {
		fmt.Println("file read error!")
	}
	return string(buf)
}

// 测试for Range
func test_for_Range() {
	ss := []string{"a", "b"}

	for k, v := range ss {
		fmt.Printf("key is %d,value is %s\n", k, v)
	}
}

func test_json_process() {
	buf, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("file read error!")
	}

	if !gjson.ValidBytes(buf) {
		fmt.Println("error jsong: ", string(buf))
		return
	}
	result := gjson.ParseBytes(buf)
	fmt.Println(result)

	id := result.Get("core.container.job.id")

	fmt.Println("id is ", id)

	res := result.Get("transformer")

	if res.Exists() {
		fmt.Println(res)
	} else {
		fmt.Println("transformer error")
	}

	var jsons []string
	switch {
	case res.IsArray():
		a := res.Array()
		fmt.Println("In switch")
		for _, v := range a {
			jsons = append(jsons, v.Str)
			fmt.Println("In loop")
		}
	}
	fmt.Println("jsons ", jsons)
	// invoke(d)
}

func str2gbk(text string) string {
	TextBuff, err := simplifiedchinese.GBK.NewEncoder().String(text)
	if err != nil {
		fmt.Println(err)
	}
	return TextBuff
}
