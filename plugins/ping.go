package plugins

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

const (
	ICMP_ECHO = 8 // 回送请求
)

// ICMP 请求报文头结构
type ICMP struct {
	Type        uint8  // 类型
	Code        uint8  // 代码
	CheckSum    uint16 // 校验和
	ID          uint16 // ID
	SequenceNum uint16 // 序号
}

var (
	size    = 32   // 数据大小
	count   = 5    // 请求数量
	timeout = 3000 // 超时时长
)

func Ping(target string) bool {
	// DNS解析
	raddr, err := net.ResolveIPAddr("ip", target)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	//logger.Info("Ping " + target + " [" + raddr.String() + "]")

	laddr := net.IPAddr{IP: net.ParseIP("0.0.0.0")}
	// 返回一个 ip socket
	conn, err := net.DialIP("ip4:icmp", &laddr, raddr)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	defer conn.Close()

	icmp := ICMP{ICMP_ECHO, 0, 0, 1, 1}
	var buffer bytes.Buffer
	err = binary.Write(&buffer, binary.BigEndian, icmp)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	// 请求的数据
	data := make([]byte, size)
	// 将请求数据写到 icmp 报文头后
	buffer.Write(data)
	data = buffer.Bytes()
	// ICMP 请求签名（校验和）：相邻两位拼接到一起，拼接成两个字节的数
	binary.BigEndian.PutUint16(data[2:], checkSum(data))

	// 设置超时时间
	err = conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	recv := make([]byte, 1024)
	aliveFlag := false

	for i := 0; i < count; i++ {
		startTime := time.Now()
		if _, err := conn.Write(data); err != nil {
			//logger.Error(err.Error())
			time.Sleep(time.Second)
			continue
		}
		length, err := conn.Read(recv)
		if err != nil {
			//logger.Error(err.Error())
			time.Sleep(time.Second)
			continue
		}
		t := time.Since(startTime).Milliseconds()
		aliveFlag = true
		logger.Debug(fmt.Sprintf("来自 %d.%d.%d.%d 的回复：字节=%d 时间=%dms TTL=%d\n", recv[12], recv[13], recv[14], recv[15], length-28, t, recv[8]))
		logger.Info(target + " alive")
		break
	}
	return aliveFlag
}

// 求校验和
func checkSum(data []byte) uint16 {
	// 第一步：两两拼接并求和
	length := len(data)
	index := 0
	var sum uint32
	for length > 1 {
		// 拼接且求和
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		length -= 2
		index += 2
	}
	// 奇数情况，还剩下一个，直接求和过去
	if length == 1 {
		sum += uint32(data[index])
	}

	// 第二部：高 16 位，低 16 位 相加，直至高 16 位为 0
	hi := sum >> 16
	for hi != 0 {
		sum = hi + uint32(uint16(sum))
		hi = sum >> 16
	}
	// 返回 sum 值 取反
	return uint16(^sum)
}
