package main

import (
	"net"
	"fmt"
	"time"
)

//  UDP编程


/*

UDP（User Datagram Protocol），UDP是无连接的传输协议
TCP：有连接的传输协议


 */

func main() {
	service := ":8888"
	updAddr,_ := net.ResolveUDPAddr("udp4",service)

	conn,_ := net.ListenUDP("udp", updAddr)
	fmt.Println("UDP时间服务器已经启动")

	var buf[128]byte
	dataLen, addr,err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	fmt.Println(string(buf[:dataLen]))
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime),addr)
}