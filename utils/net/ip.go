package net

// 获取本机IP地址
import "net"

var Ip string
var Ips []string

func init() {
	Ips, _ = LocalIPv4s()
	Ip = Ips[0]
}

func LocalIPv4s() ([]string, error) {
	var ips []string
	address, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range address {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}
