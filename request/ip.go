package request

import (
	"github.com/thinkeridea/go-extend/exnet"
	"net"
	"net/http"
	"strings"
)

//获取IP
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := exnet.ClientPublicIP(req); ip != "" {
		remoteAddr = ip
	} else if ip := exnet.ClientIP(req); ip != "" {
		remoteAddr = ip
	} else if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

//获取当前服务器IP
func CurNodeIp() string {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		return err.Error()
	}
	defer conn.Close()
	nodeip := strings.Split(conn.LocalAddr().String(), ":")[0]
	return nodeip
}