package main

import (
	"net/http"
	"fmt"
	"log"
	"net/url"
)

/*
处理同名请求字段


获取请求字段值有2种方式

1.  r.Form
返回一个url.Values类型的值，如果字段不存在，返回一个长度为0的集合

2.  r.Form.Get

只返回第一个叫fieldname的字段值，如果字段不存在，返回一个空串（长度为0的字符串）

 */

func sameNameEcho(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()   //  分析客户端的body数据
	fmt.Println(r.Form)

	usernames,ok := r.Form["username"]
	if ok == false {
		return
	}

	for _,username := range usernames {
		fmt.Println(username)
	}

	fmt.Println("用户名：" + r.Form.Get("username"))

	r.Form.Set("username","李宁")
	fmt.Println(r.Form)


}
func main() {

	// 测试url.Values
	v := url.Values{}
	v.Add("friend","John")
	v.Add("friend","Mike")
	v.Add("friend","马云")

	fmt.Println(v.Encode())

	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])

	// 指定路由和回调函数
	http.HandleFunc("/",sameNameEcho)
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}