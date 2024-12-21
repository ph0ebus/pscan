package controller

import (
	"encoding/json"
	"net/http"
	"pscan/plugins"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Params struct {
	IpRange   string
	ScanType  string
	PortRange string
}

type WebSocketConnection struct {
	Conn *websocket.Conn
	Lock sync.Mutex
}

func Scan(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	conn := &WebSocketConnection{
		Conn: ws,
		Lock: sync.Mutex{},
	}

	for {
		mt, message, _ := ws.ReadMessage()
		var params Params
		json.Unmarshal(message, &params)
		if params.IpRange == "" {
			ws.WriteMessage(mt, []byte("[!] Invalide IP Range, please input a valid IP range"))
			return
		}
		if params.ScanType == "" {
			ws.WriteMessage(mt, []byte("[!] please choice a valid scan type"))
			return
		}

		ips := plugins.ParseIPs(params.IpRange)
		aliveIps := []string{}
		wg := &sync.WaitGroup{}
		for _, ip := range ips {
			wg.Add(1)
			go func() {
				if plugins.Ping(ip) {
					conn.Lock.Lock()
					ws.WriteMessage(mt, []byte("[+] "+ip+" alive"))
					aliveIps = append(aliveIps, ip)
					conn.Lock.Unlock()
				}
				wg.Done()
			}()
		}
		wg.Wait()

		if params.ScanType == "host" {

		} else {
			ports := plugins.ParsePorts(params.PortRange)
			for _, ip := range aliveIps {
				openPorts := []int{}

				wg := &sync.WaitGroup{}
				for _, port := range ports {
					wg.Add(1)
					go func() {
						if plugins.PortConnect(ip, port) {
							conn.Lock.Lock()
							openPorts = append(openPorts, port)
							if params.ScanType == "service" {
								serviceDescription := plugins.ServiceDetect(ip, port)
								ws.WriteMessage(mt, []byte("[+] Service Scan "+ip+" "+strconv.Itoa(port)+" maybe "+serviceDescription))
							} else {
								ws.WriteMessage(mt, []byte("[+] "+ip+" "+strconv.Itoa(port)+" open"))
							}
							conn.Lock.Unlock()
						}
						wg.Done()
					}()
				}
				wg.Wait()
			}
		}

		ws.WriteMessage(mt, []byte("[+] Done"))
	}
}
