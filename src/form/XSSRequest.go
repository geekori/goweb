package main

import (
	"net/http"
	"fmt"
	"log"
	"text/template"
	"regexp"
)

/*

防止跨站脚本攻击（XSS）Cross Site Scripting


Web页面：

1. 静态
2. 动态，是由服务端产生的

（1）完全由服务端自己产生的
（2）动态内容的全部或部分是由用户提交的


 */


func xssEcho(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	code := r.Form.Get("code")
	fmt.Println(code)
    code = template.HTMLEscapeString(code)
    fmt.Println(code)
	//fmt.Fprintf(w,"<html>" + code + "</html>")
	t,_ := template.New("test").Parse(`<html> {{ . }}</html>`)
	t.ExecuteTemplate(w,"test",code)

}
func xssEcho1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	code := r.Form.Get("code")
	fmt.Println(code)
	reg,_ := regexp.Compile(`<script[^>]*>|</script>`)
	text := reg.ReplaceAllLiteralString(code,"")
	//fmt.Fprintf(w,"<html>" + code + "</html>")
	t,_ := template.New("test").Parse(`<html> {{ . }}</html>`)
	t.ExecuteTemplate(w,"test",text)

}
func main() {
	// 指定路由和回调函数
	http.HandleFunc("/xss",xssEcho)
	http.HandleFunc("/xss1",xssEcho1)
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}
