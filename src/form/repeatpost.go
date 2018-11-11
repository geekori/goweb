package main

import (
	"net/http"
	"time"
	"fmt"
	"crypto/sha512"
	"io"
	"strconv"
	"html/template"
	"log"
)

/*

防止重复提交表单

在Web页面放置一个隐藏的字段（隐藏的<input>组件），在页面每次装载后，服务端都会为页面生成一个
唯一的Token。如果只是通过回退（back）显示原来的页面，那么Token不会变化

如果发现Token重复提交，那么就不会处理提交上来的数据，会直接报错


 */

func showRepeatPostPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 核心代码需要产生Token

		currentTime := time.Now().Unix()
		fmt.Println(currentTime)

		h := sha512.New()
		io.WriteString(h,strconv.FormatInt(currentTime,10))
		//  产生Token
		token := fmt.Sprintf("%x",h.Sum(nil))
		fmt.Println(token)

		//  通过模板装载页面
		t,_ := template.ParseFiles("./src/form/repeatpost.html")
		t.Execute(w,token)


	}
}
var count = 0
func repeatPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		return
	}

	r.ParseForm()
	token := r.Form.Get("token")
	if token != "" {
		fmt.Println("校验Token，持有该Token的页面只能提交一次")
		if count > 0 {
			fmt.Fprintln(w,"您重复提交了")
			return
		}
		fmt.Println(count)
		count++
	} else {
		fmt.Fprintf(w,"Token不存在")
	}
	name := r.Form.Get("name")
	fmt.Fprintf(w,name)

}
func main() {
	// 指定路由和回调函数
	http.HandleFunc("/",showRepeatPostPage)
	http.HandleFunc("/repeatpost",repeatPost)
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}