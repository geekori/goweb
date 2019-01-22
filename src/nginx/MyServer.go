package main

import (
	"net/http"
	"fmt"
	"log"
)

/*
Go Web与Nginx集成

Nginx：使用C语言编写的高性能Web服务器

Nginx尤其擅长处理静态资源（html、js、css、image）


brew install nginx

 */

 func login(w http.ResponseWriter,r *http.Request) {
 	r.ParseForm()
 	username := r.Form["username"][0]
 	password := r.Form["password"][0]

 	if username == "john" && password=="4321" {
 		fmt.Fprintf(w,"success")
	} else {
		fmt.Fprintf(w,"failed")
	}
 }

 func main() {
 	http.HandleFunc("/mylogin",login)

	 fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
	 //  启动HTTP服务，并监听端口号
	 err := http.ListenAndServe(":8900",nil)

	 if err != nil {
		 log.Fatal("ListenAndServe:",err)
	 }
 }