package main

import (
	"net/http"
	"fmt"
	"log"
)

/*
用表单提交用户登录信息（POST请求）
 */

func postEcho(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()   //  分析客户端的body数据
	fmt.Println(r.Form)

	usernames,ok := r.Form["username"]
	if ok == false {
		return
	}
	username := usernames[0]
	password := r.Form["password"][0]

	if username == "Bill" && password == "1234" {
		fmt.Fprintf(w, "登录成功")
	} else {
		fmt.Fprintf(w, "登录失败")
	}

}
func main() {
	// 指定路由和回调函数
	http.HandleFunc("/",postEcho)
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}