package main

import (
	"fmt"
	"net"
	"time"
)

// shuru:   telnet localhost 8088
func main() {
	service := ":8080"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)
	fmt.Println("time server is ready")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		fmt.Println(daytime)
		conn.Close()
	}
}
