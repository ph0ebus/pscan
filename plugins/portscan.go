package plugins

import (
	"net"
	"strconv"
	"time"
)

func PortConnect(ip string, port int) bool {
	target := ip + ":" + strconv.Itoa(port)
	//tcpAddr, _ := net.ResolveTCPAddr("tcp", target)

	// 使用DialTimeout设置两秒超时
	conn, err := net.DialTimeout("tcp", target, 2*time.Second)
	if err != nil {
		return false
	}
	logger.Info(target + " open")
	defer conn.Close()
	return true
}
