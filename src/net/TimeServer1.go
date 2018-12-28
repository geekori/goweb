package main

import (
	"net"
	"fmt"
	"time"
)

//  支持同时处理多个客户端请求的时间服务器


func main() {
	tcpAddr,_ := net.ResolveTCPAddr("tcp4",":9999")

	listener,_ := net.ListenTCP("tcp", tcpAddr)

	fmt.Println("时间服务器已经启动...")
	for {
		conn,err := listener.Accept() // 将会被阻塞，直到收到客户端请求为止
		if err != nil {
			continue
		}
		go handleClient1(conn)  // 异步执行
	}



}

func handleClient1(conn net.Conn) {
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
	conn.Close()
}


