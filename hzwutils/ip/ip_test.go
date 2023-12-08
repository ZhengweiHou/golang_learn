package ip

import (
	"fmt"
	"net"
	"regexp"
	"testing"
)

func TestIp(t *testing.T) {
	ip := GetLocalIP()
	fmt.Println(ip)

	// iprg1 := NewIPRange("172.17", "172.17")
	// fmt.Println(iprg1.IsEnabled("172.17.0.0"))

	InitIPRanges("172.17:172.17")
	ip = GetLocalIP()
	fmt.Println(ip)

	InitIPRanges("192.168.105")
	ip = GetLocalIP()
	fmt.Println(ip)

	oriport := 8888
	address := fmt.Sprintf("%s:%d", ip, oriport)
	listener, _ := net.Listen("tcp", address)
	defer listener.Close()
	port, _ := GetAvailablePort(ip, 8888)
	fmt.Println(port) // 8889

}

func TestIpRangeRegex(t *testing.T) {
	// 测试字符串
	testStrings := []string{
		"10",
		"10.250",
		"10.250:10.252",
		"10,12",
		"127.0.0.1:127.0.1.0,10.250:10.252",
		"127.0.0.1",
		"10.250.0.1",
		"10.255",
		"10.250.10:10.230,10.233",
		"0:255",
		"0:255,0:255",
		// false 类型
		"abc",
		"0:256,0:255",
	}

	// var ipRangeRegex *regexp.Regexp = regexp.MustCompile(`^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d))*$`)
	// var ipRangeRegex *regexp.Regexp = regexp.MustCompile(`^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){0,3}$`)
	// var ipRangeRegex *regexp.Regexp = regexp.MustCompile(`^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){0,3}(:(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){0,3}){0,1}$`)

	// var ipRangeRegex *regexp.Regexp = regexp.MustCompile(`^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){0,3}(:(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){0,3}){0,1}(,(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){0,3}(:(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){0,3}){0,1})*$`)

	reg := `(2[0-4]\d|25[0-5]|[01]?\d\d?)(\.(2[0-4]\d|25[0-5]|[01]?\d\d?)){0,3}`
	reg = fmt.Sprintf(`(%s)(:(%s)){0,1}`, reg, reg)
	reg = fmt.Sprintf(`^(%s){1}(,(%s){0,3})*$`, reg, reg)
	var ipRangeRegex *regexp.Regexp = regexp.MustCompile(reg)
	fmt.Println(reg)

	// 检查每个字符串是否匹配正则表达式
	for _, str := range testStrings {
		match := ipRangeRegex.MatchString(str)
		fmt.Printf("%s: %t\n", str, match)
	}
}

func TestIpRangeRegex2(t *testing.T) {
	testStrings := []string{
		"10",
		"127.0.0.1",
		"10.250.0.1",
		"1.1.1.1",
		"0.0.0.0",
		"0.1.0.0",
	}
	var ipRangeRegex *regexp.Regexp = regexp.MustCompile(`^(2[0-4]\d|25[0-5]|[01]?\d\d?)(\.(2[0-4]\d|25[0-5]|[01]?\d\d?)){3}$`)
	// var ipRangeRegex *regexp.Regexp = regexp.MustCompile(`^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){3}$`)
	// 检查每个字符串是否匹配正则表达式
	for _, str := range testStrings {
		match := ipRangeRegex.MatchString(str)
		fmt.Printf("%s: %t\n", str, match)
	}
}
