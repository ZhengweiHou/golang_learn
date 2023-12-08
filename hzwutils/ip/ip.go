package ip

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

const (
	minPort int = 0
	maxPort int = 65535
)

// IPRanges ip范围，用于获取本地ip时指定范围
var IPRanges []IPRange

// enableIPRange 的格式检查：,号分割多个配置，:号分割开始和结束ip段，.号分割的是ipv4的段，总ipv4的段数可以是1~4个，
// 案例：10.250.10:10.230,10.233、10.250
// var ipRangeRegex *regexp.Regexp = regexp.MustCompile(`^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){0,3}(:(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){0,3}){0,1}(,(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){0,3}(:(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])(\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)){0,3}){0,1})*$`)
var ipRangeRegex *regexp.Regexp = regexp.MustCompile(`^(((2[0-4]\d|25[0-5]|[01]?\d\d?)(\.(2[0-4]\d|25[0-5]|[01]?\d\d?)){0,3})(:((2[0-4]\d|25[0-5]|[01]?\d\d?)(\.(2[0-4]\d|25[0-5]|[01]?\d\d?)){0,3})){0,1}){1}(,(((2[0-4]\d|25[0-5]|[01]?\d\d?)(\.(2[0-4]\d|25[0-5]|[01]?\d\d?)){0,3})(:((2[0-4]\d|25[0-5]|[01]?\d\d?)(\.(2[0-4]\d|25[0-5]|[01]?\d\d?)){0,3})){0,1}){0,3})*$$`)

// InitIPRanges 初始化IPRange列表
func InitIPRanges(enableIPRange string) {
	if len(enableIPRange) != 0 && !ipRangeRegex.MatchString(enableIPRange) {
		// TODO 日志警告，enableIPRange配置非法
		return
	}

	IPRanges = make([]IPRange, 0)
	if len(enableIPRange) == 0 {
		IPRanges = append(IPRanges, NewIPRange("0", "255"))
	} else {
		ipRanges := strings.Split(enableIPRange, ",")
		for _, ipRange := range ipRanges {
			if len(ipRange) == 0 {
				continue
			}
			if strings.Contains(ipRange, ":") {
				ranges := strings.Split(ipRange, ":")
				IPRanges = append(IPRanges, NewIPRange(ranges[0], ranges[1]))
			} else {
				IPRanges = append(IPRanges, NewIPRange(ipRange, ipRange))
			}
		}
	}
}

/*
GetLocalIP 获得本地的网络地址
在有超过一块网卡时有问题，这里每次只取了第一块网卡绑定的IP地址 当存在这种情况的时候，就需要InitIPRanges初始化IPRanges，用以限制IP范围

@return string 返回本地ip，失败则返回空串
*/
func GetLocalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, inter := range inters {
		// 排除回环网口
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						// ip 范围可用判断
						if IPIsRangeAvailable(ipnet.IP.String()) {
							return ipnet.IP.String()
						}

					}
				}
			}
		}
	}
	return "0.0.0.0"
}

// IPIsRangeAvailable ip范围可用判断
func IPIsRangeAvailable(ip string) bool {
	if len(ip) == 0 {
		return false
	}

	// IPRange 为空则返回true
	if IPRanges == nil || len(IPRanges) < 1 {
		return true
	}

	// 当前ip是否在配置的ip范围内
	for _, iprang := range IPRanges {
		if iprang.IsEnabled(ip) {
			return true
		}
	}

	return false
}

// GetAvailablePort 检查当前指定端口是否可用，不可用则自动+1再试（随机端口从默认端口开始检查）
func GetAvailablePort(host string, oriPort int) (int, error) {

	// 检查host是否可用的
	if !IsHostAvailable(host) {
		return 0, fmt.Errorf("ERROR_HOST_NOT_FOUND")
	}

	port := oriPort
	if port < minPort {
		port = minPort
	}

	for port < maxPort {
		if isPortAvailable(host, port) {
			return port, nil
		}
		port++
	}

	return 0, fmt.Errorf("ERROR_BIND_PORT_ERROR")
}

// isPortAvailable 检查指定主机上的给定端口是否可用
func isPortAvailable(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return false
	}
	defer listener.Close()
	return true
}

// IsHostAvailable 检查host是否可用
func IsHostAvailable(host string) bool {
	// isAnyHost
	if strings.EqualFold("0.0.0.0", host) {
		return true
	}
	// isLocalHost
	if strings.EqualFold("127.0.0.1", host) || strings.EqualFold("localhost", host) {
		return true
	}

	// isHostInNetWork 检查
	return isHostInNetworkCard(host)
}

// isHostInNetworkCard 是否网卡上的地址
func isHostInNetworkCard(host string) bool {
	addr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		return false
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return false
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok {
				if ipnet.IP.Equal(addr.IP) {
					return true
				}
			}
		}
	}

	return false
}

// IPRange ip范围
type IPRange struct {
	Start int64
	End   int64
}

// NewIPRange 构建IPRange
func NewIPRange(startIP, endIP string) IPRange {
	return IPRange{
		Start: parseStart(startIP),
		End:   parseEnd(endIP),
	}
}

func parseStart(ip string) int64 {
	segments := []int{0, 0, 0, 0}
	return parse(segments, ip)
}

func parseEnd(ip string) int64 {
	segments := []int{255, 255, 255, 255}
	return parse(segments, ip)
}

func parse(segments []int, ip string) int64 {
	ipSegments := strings.Split(ip, ".")
	for i := 0; i < len(ipSegments); i++ {
		segments[i], _ = strconv.Atoi(ipSegments[i])
	}
	var ret int64
	for i := 0; i < len(segments); i++ {
		ret = ret*256 + int64(segments[i])
	}
	return ret
}

// IsEnabled 判断指定的ip是否在IP范围内
func (r IPRange) IsEnabled(ip string) bool {
	ipSegments := strings.Split(ip, ".")
	var ipInt int64
	for _, ipSegment := range ipSegments {
		val, _ := strconv.Atoi(ipSegment)
		ipInt = ipInt*256 + int64(val)
	}
	return ipInt >= r.Start && ipInt <= r.End
}
