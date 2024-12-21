package plugins

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

func ParseIPs(target string) (hosts []string) {
	if strings.Contains(target, ",") {
		lists := strings.Split(target, ",")
		for _, list := range lists {
			hosts = append(hosts, parseIP(list)...)
		}
	} else {
		hosts = parseIP(target)
	}

	return hosts
}

func parseIP(target string) []string {
	matched, _ := regexp.MatchString("[a-zA-Z]+", target)
	if matched {
		// domain
		_, err := net.ResolveIPAddr("ip", target)
		if err != nil {
			logger.Error(err.Error())
			return []string{}
		}
		return []string{target}
	} else if strings.Contains(target, "/") {
		//  192.168.1.1/24
		return parseCIDR(target)
	} else if strings.Contains(target, "-") {
		//  192.168.1.1-255
		//  192.168.1.1-192.168.2.1
		return parseRangeIP(target)
	} else {
		// single ip
		return []string{target}
	}
}

func parseCIDR(target string) (hosts []string) {
	_, ipNet, err := net.ParseCIDR(target)
	if err != nil {
		logger.Error(err.Error())
		return []string{}
	}

	for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		hosts = append(hosts, ip.String())
	}

	return removeHostsDuplicate(hosts)
}

func parseRangeIP(target string) (hosts []string) {
	parts := strings.Split(target, "-")
	if len(parts) != 2 {
		logger.Error(target + " is not a valid range")
		return []string{}
	}

	startIP := net.ParseIP(parts[0])
	var endIP net.IP
	if strings.Contains(parts[1], ".") {
		endIP = net.ParseIP(parts[1])
	} else {
		startIP4 := startIP.To4()
		end, err := strconv.Atoi(parts[1])
		if err != nil || end > 256 || end < int(startIP4[3]) {
			logger.Error(target + " is not a valid IP range")
			return []string{}
		}
		endIP4 := make(net.IP, len(startIP4))
		copy(endIP4, startIP4)
		endIP4[3] = byte(end)
		endIP = endIP4.To16()
	}

	for ip := startIP; !ip.Equal(endIP); inc(ip) {
		hosts = append(hosts, ip.String())
	}
	hosts = append(hosts, endIP.String())

	return hosts
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] != 0 {
			break
		}
	}
}

// 去重
func removeHostsDuplicate(old []string) []string {
	var result []string
	temp := map[string]struct{}{}
	for _, item := range old {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
