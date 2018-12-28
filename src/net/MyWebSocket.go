package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"net/http"
)

/*
WebSocket实战

Socket：TCP

Web：可以用于Web程序

WebSocket看做是在Web前端页面中使用的Socket

在WebSocket出现之前，Web页面主要通过轮询的方式从服务端获取数据

基于TCP

TCP：HTTP、HTTPS、FTP、SMTP、POP3、WebSocket

ws://ip   HTTP
wss://ip  HTTPS  SSL

go get golang.org/x/net/websocket   VPN

GOPATH
 */

 // WebSocket

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws,&reply);err != nil {
			fmt.Println("不能接收数据")
			break
		}

		fmt.Println("来至客户端的数据：" + reply)

		msg := "返回给客户端的数据：" + reply

		fmt.Println("正在发送数据给客户端" + msg)

		if err = websocket.Message.Send(ws,msg);err != nil {
			fmt.Println("发送数据失败")
			break
		}
	}
}

func Echo1(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws,&reply);err != nil {
			fmt.Println("不能接收数据")
			break
		}

		fmt.Println("来至客户端的数据：" + reply)

		msg := reply

		fmt.Println("正在发送数据给客户端" + msg)

		if err = websocket.Message.Send(ws,msg);err != nil {
			fmt.Println("发送数据失败")
			break
		}
	}
}
func main() {
	http.Handle("/",websocket.Handler(Echo))
	http.Handle("/abc",websocket.Handler(Echo1))
	fmt.Println("WebSocket服务器已经启动")
	// ws://127.0.0.1
	if err := http.ListenAndServe(":1234",nil); err != nil {
		fmt.Println("监听错误")
	}
}