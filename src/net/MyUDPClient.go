package main

import (
	"net"
	"fmt"
)

// UDP客户端

func main() {
	udpAddr,_ := net.ResolveUDPAddr("udp4","127.0.0.1:8888")
	conn,_ := net.DialUDP("udp",nil,udpAddr)

	conn.Write([]byte("hello world"))

	var buf [512]byte
	n,_ := conn.Read(buf[0:])
	fmt.Println(string(buf[0:n]))
}