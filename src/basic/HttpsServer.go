package main

import (
	"net/http"
	"fmt"
	"log"
)

/*


net/http

func funname(w http.ResponseWriter, r *http.Request) {

}

 */

func echo1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path:", r.URL.Path)
	fmt.Fprintf(w,"Hello Go Web")
}

func main() {
    // 指定路由和回调函数
    http.HandleFunc("/",echo2)
    fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
    //  启动HTTP服务，并监听端口号
    err := http.ListenAndServe(":8900",nil)
    fmt.Println("监听之后")
    if err != nil {
    	log.Fatal("ListenAndServe:",err)
	}
}

