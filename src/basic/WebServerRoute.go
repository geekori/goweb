package main

import (
	"net/http"
	"fmt"
	"log"
)

/*
设置路由


 */

func echo2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path:", r.URL.Path)
	fmt.Fprintf(w,"Hello Go Web")
}
func echo21(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path:", r.URL.Path)
	fmt.Fprintf(w,"Hello World")
}
func main() {
    // 指定路由和回调函数
    http.HandleFunc("/",echo2)  // http://localhost:8900/a/b
    http.HandleFunc("/a/b/",echo21) //http://localhost:8900/a/b/c/d
    fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
    //  启动HTTP服务，并监听端口号
    err := http.ListenAndServe(":8900",nil)
    fmt.Println("监听之后")
    if err != nil {
    	log.Fatal("ListenAndServe:",err)
	}
}

