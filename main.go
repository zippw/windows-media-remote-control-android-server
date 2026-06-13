package main

import (
	"fmt"
	"net"
	"net/http"
	"pcrc/api"
	"strings"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil { return "127.0.0.1" }
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipStr := ipnet.IP.String()
				if !strings.HasPrefix(ipStr, "172.") && !strings.HasPrefix(ipStr, "169.254") {
					return ipStr
				}
			}
		}
	}
	return "127.0.0.1"
}

func main() {
	const port = "12597"

	api.RegisterRoutes()

	localIP := getLocalIP()

	fmt.Printf("Local server: http://%s:%s\n", localIP, port)
	
	var err = http.ListenAndServe(":"+port, nil)
	if err != nil { fmt.Printf("ERROR: %v\n", err) }
}