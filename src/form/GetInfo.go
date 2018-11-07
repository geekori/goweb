package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

/*
处理表单：获取HTTP Get请求字段值

HTTP Get

HTTP Post


 */

func echo1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	names,ok := r.Form["name"]
	if ok == true {
		fmt.Println(names[0])
		fmt.Println(r.Form["age"][0])
	}
	for k,v := range r.Form {
		fmt.Println("key:",k)
		fmt.Println(v)
		fmt.Println("val:", strings.Join(v,""))

	}
	fmt.Fprintf(w,"Hello Get")
}



func main() {
	// 指定路由和回调函数
	http.HandleFunc("/",echo1)
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}