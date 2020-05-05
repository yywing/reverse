package utils

import (
	"fmt"
	"net"
)

// IsIPv6 不支持 "IPv4映像地址".
func IsIPv6(raw string) bool {
	if ip := net.ParseIP(raw); ip != nil {
		return ip.To4() == nil
	}
	return false
}

func IPAndPort(ip string, port uint16) string {
	if IsIPv6(ip) {
		return "[" + ip + "]" + ":" + fmt.Sprintf("%v", port)
	}
	return ip + ":" + fmt.Sprintf("%v", port)
}
