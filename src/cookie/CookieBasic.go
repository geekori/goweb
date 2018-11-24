package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
)

/*

读写Cookie


当客户端第一次访问服务端时，服务端为客户端发一个凭证（全局唯一的字符串）

服务端将字符串发送给客户端（写Cookie的过程）

HTTP请求头（发送给服务端Cookie）

HTTP响应头（在服务端通知客户端保存Cookie）


 */

func writeCookie(w http.ResponseWriter, r *http.Request)  {
    expiration := time.Now()
    expiration = expiration.AddDate(0,0,3)

	cookie := http.Cookie{Name:"username", Value:"geekori",Expires:expiration}

	http.SetCookie(w,&cookie)
	fmt.Fprintf(w,"write cookie success")
}
func readCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"<html>")
	cookie,_ := r.Cookie("username")
	fmt.Fprint(w,cookie)
	fmt.Fprint(w,"<br>")
	fmt.Fprint(w,cookie.Value)
	fmt.Fprint(w,"</html>")
}
func main() {
	// 指定路由和回调函数
	http.HandleFunc("/writeCookie",writeCookie)
	http.HandleFunc("/readCookie",readCookie)
	fmt.Println("服务器已经启动")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}


