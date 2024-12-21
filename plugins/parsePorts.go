package plugins

import (
	"strconv"
	"strings"
)

func ParsePorts(p string) (ports []int) {
	if strings.Contains(p, ",") {
		lists := strings.Split(p, ",")
		for _, list := range lists {
			ports = append(ports, parsePort(list)...)
		}
	} else {
		ports = parsePort(p)
	}

	return removePortsDuplicate(ports)
}

func parsePort(p string) (ports []int) {
	if strings.Contains(p, "-") {
		ports = parseRangePort(p)
	} else {
		port, err := strconv.Atoi(p)
		if err != nil {
			logger.Error(err.Error())
		}
		ports = []int{port}
	}
	return ports
}

func parseRangePort(p string) (ports []int) {
	parts := strings.Split(p, "-")
	startPort, err := strconv.Atoi(parts[0])
	if err != nil || startPort > 65535 || startPort < 1 {
		logger.Error(p + " is not a valid IP range")
	}
	endPort, err := strconv.Atoi(parts[1])
	if err != nil || endPort > 65535 || endPort < startPort {
		logger.Error(p + " is not a valid IP range")
	}

	for i := startPort; i <= endPort; i++ {
		ports = append(ports, i)
	}
	return ports
}

// 去重
func removePortsDuplicate(old []int) []int {
	var result []int
	temp := map[int]struct{}{}
	for _, item := range old {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
