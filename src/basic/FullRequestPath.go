package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
)

/*
获取完整的请求路径

http://localhost:8900/a/b/x.html

1. scheme：http/https
ok. 2. 域名或IP：localhost
ok. 3. 端口号：8900
ok. 4. Path：/a/b/x.html

 */

func echo4(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println("scheme：" + r.URL.Scheme)
    fmt.Println("域名（IP）和端口号：" + r.Host)
	fmt.Println("Path:" + r.RequestURI)
    scheme := "http://"
    if r.TLS != nil {
    	scheme = "https://"
	}
	fmt.Println("完整的请求路径：" + strings.Join([]string{scheme,r.Host,r.RequestURI},""))
    fmt.Fprintf(w,"Hello Go Web")

}

func main() {
    // 指定路由和回调函数
    http.HandleFunc("/",echo4)
    fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
    //  启动HTTP服务，并监听端口号
    err := http.ListenAndServe(":8900",nil)
    fmt.Println("监听之后")
    if err != nil {
    	log.Fatal("ListenAndServe:",err)
	}
}

