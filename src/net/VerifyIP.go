package main

import (
	"net"
	"fmt"
	"reflect"
)

/*

Socket

HTTP、FTP、SMTP、POP3

IP：IPV4   IPV6

xxx.xxx.xxx.xxx

xxx：0到255

192.168.1.1

 */

func main() {
	addr := net.ParseIP("192.168.1.1")  // byte[]
	fmt.Println(reflect.TypeOf(addr))
	if addr == nil {
		fmt.Println("不正确的IP地址")
	} else {
		fmt.Println("IP地址：", addr.String())
	}
}


