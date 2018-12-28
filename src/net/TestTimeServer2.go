package main

import (
	"net"
	"io/ioutil"
	"fmt"
)

//  用于测试时间服务器

func main() {

	tcpAddr,_ := net.ResolveTCPAddr("tcp4","127.0.0.1:9999")
	//  连接服务器
	conn,_:= net.DialTCP("tcp",nil,tcpAddr)

	var data string
	data = "timestamp1"

	conn.Write([]byte(data))




	result,_ := ioutil.ReadAll(conn)
	fmt.Println(string(result))


}
