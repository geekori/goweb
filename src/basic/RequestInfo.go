package main

import (
	"net/http"
	"fmt"
	"log"
)

/*
获取HTTP请求头信息

1. HTTP请求头：Path、Host、Method（Get、Post）、Proto、UserAgent
2. Body

 */

func echo3(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Path:",r.URL.Path)
    fmt.Println("Url:",r.URL)   // /a/b/
    fmt.Println("Host",r.Host)  // localhost:8900
    fmt.Println("Header:", r.Header)
    fmt.Println("Method:",r.Method)
    fmt.Println("Proto：", r.Proto)

    fmt.Println("UserAgent:", r.UserAgent())
    fmt.Fprintf(w,"Hello Go Web")

}

func main() {
    // 指定路由和回调函数
    http.HandleFunc("/",echo3)
    fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
    //  启动HTTP服务，并监听端口号
    err := http.ListenAndServe(":8900",nil)
    fmt.Println("监听之后")
    if err != nil {
    	log.Fatal("ListenAndServe:",err)
	}
}

