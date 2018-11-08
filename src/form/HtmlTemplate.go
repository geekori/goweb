package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)

/*
使用HTML模板显示登录页面


模板：静态+动态
 */

func login(w http.ResponseWriter, r *http.Request) {
	//  保证只能处理POST请求
	if r.Method == "GET" {
		return
	}
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
//  用于显示静态页面（登录页面）
func showLoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t,_:=template.ParseFiles("/MyStudio/video_src/goweb/src/form/login.tpl")
		t.Execute(w,nil)
	}
}

func main() {
	// 指定路由和回调函数
	http.HandleFunc("/",showLoginPage)
	http.HandleFunc("/login",login)
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}