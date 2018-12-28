package main

import (
	"net"
	"fmt"
	"time"
	"strings"
	"strconv"
)

//  可以接收客户端数据的时间服务器

func main() {
	tcpAddr,_ := net.ResolveTCPAddr("tcp4",":9999")

	listener,_ := net.ListenTCP("tcp", tcpAddr)

	fmt.Println("时间服务器已经启动...")
	for {
		conn,err := listener.Accept() // 将会被阻塞，直到收到客户端请求为止
		if err != nil {
			continue
		}
		go handleClient2(conn)  // 异步执行
	}



}

func handleClient2(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))

	request := make([]byte,1024)

	for {
		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if readLen == 0 {
			break
		} else if strings.TrimSpace(string(request[:readLen])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(),10)
			conn.Write([]byte(daytime))
			conn.Close()
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
			conn.Close()
		}

	}

}


