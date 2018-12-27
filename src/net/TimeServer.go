package main

import (
	"net"
	"fmt"
	"os"
	"time"
)

/*
实现基于Socket的时间服务器

 */

func main() {
	tcpAddr,err := net.ResolveTCPAddr("tcp4",":9999")
	CheckError(err)
	listener,err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)
	fmt.Println("时间服务器已经启动...")
	for {
		conn,err := listener.Accept() // 将会被阻塞，直到收到客户端请求为止
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}



}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
		os.Exit(1)
	}
}

