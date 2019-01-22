package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
)

/*
nginx与包含多个路由的Go Web服务器集成

 */

 func login1(w http.ResponseWriter,r *http.Request) {
 	r.ParseForm()
 	username := r.Form["username"][0]
 	password := r.Form["password"][0]

 	if username == "john" && password=="4321" {
 		fmt.Fprintf(w,"success")
	} else {
		fmt.Fprintf(w,"failed")
	}
 }
func login2(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]

	if username == "mike" && password=="1234" {
		fmt.Fprintf(w,"success")
	} else {
		fmt.Fprintf(w,"failed")
	}
}
func login3(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]

	if username == "bill" && password=="1111" {
		fmt.Fprintf(w,"success")
	} else {
		fmt.Fprintf(w,"failed")
	}
}
//  用AJAX设置页面中的动态部分

func data(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,`{"user":"Username","pw":"Password"}`)
}
//  实现部分动态页面
// iframe
func showLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/html;charset=utf-8");
	t,_ := template.ParseFiles("/MyStudio/video_src/goweb/src/nginx/login.tpl")
	type Data struct {
		User string
		PW string
	}
	data := Data{"Username","密码"}
	t.Execute(w,data)
}
 func main() {
 	http.HandleFunc("/service/testlogin1",login1)
 	http.HandleFunc("/service/mylogin2",login2)
 	http.HandleFunc("/service/mylogin3",login3)
	 http.HandleFunc("/service/data",data)
	 http.HandleFunc("/service/showLoginPage",showLoginPage)

	 fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
	 //  启动HTTP服务，并监听端口号
	 err := http.ListenAndServe(":8900",nil)

	 if err != nil {
		 log.Fatal("ListenAndServe:",err)
	 }
 }