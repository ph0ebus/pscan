package plugins

import (
	"strconv"
)

var portServiceMap = map[int]string{
	80:    "HTTP - HyperText Transfer Protocol",
	443:   "HTTPS - Secure HyperText Transfer Protocol",
	21:    "FTP - File Transfer Protocol",
	22:    "SSH - Secure Shell",
	23:    "Telnet - Unsecured Text Communication",
	25:    "SMTP - Simple Mail Transfer Protocol",
	53:    "DNS - Domain Name System",
	110:   "POP3 - Post Office Protocol",
	143:   "IMAP - Internet Message Access Protocol",
	3306:  "MySQL - Database Service",
	5432:  "PostgreSQL - Database Service",
	6379:  "Redis - In-memory Data Structure Store",
	27017: "MongoDB - NoSQL Database",
	8080:  "HTTP Proxy or Alternative HTTP",
	69:    "TFTP - Trivial File Transfer Protocol",
	161:   "SNMP - Simple Network Management Protocol",
	162:   "SNMP Trap - Simple Network Management Protocol Trap",
	514:   "Syslog - System Logging Protocol",
	5900:  "VNC - Virtual Network Computing",
	8088:  "HTTP Proxy or Alternative HTTP",
	8888:  "HTTP Proxy or Alternative HTTP",
	10000: "Webmin - Web-based system administration",
	2049:  "NFS - Network File System",
	3389:  "RDP - Remote Desktop Protocol",
	6660:  "IRC - Internet Relay Chat",
	11211: "Memcached - Distributed Memory Object Caching",
	8181:  "HTTP Proxy or Alternative HTTP",
	9090:  "HTTP Proxy or Alternative HTTP",
	20000: "Webmin - Web-based system administration",
	4567:  "Asterisk - VoIP server",
	50000: "Oracle DB - Oracle Database Listener",
	54321: "Back Orifice - Remote Access Tool",
	9000:  "HTTP Proxy or Alternative HTTP",
	9100:  "JetDirect - Printer Service",
	5000:  "UPnP - Universal Plug and Play",
	8889:  "HTTP Proxy or Alternative HTTP",
}

func ServiceDetect(ip string, port int) string {
	serviceDescription, exists := portServiceMap[port]
	if !exists {
		serviceDescription = "Unknown Service"
	} else {
		logger.Info("Service Scan " + ip + " " + strconv.Itoa(port) + " " + serviceDescription)
	}
	return serviceDescription
}
