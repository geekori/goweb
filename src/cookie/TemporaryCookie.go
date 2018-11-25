package main

import (
	"fmt"
	"encoding/base64"
	"net/http"
	"time"
	"log"
)

/*

临时Cookie，也就是未设置过期时间的Cookie

当浏览器关闭后，Cookie就会失效

 */
func writeTemporaryCNCookies(w http.ResponseWriter, r *http.Request)  {

	var cookies map[string]string
	cookies = make(map[string]string)
	cookies["tempname"] = "超人"
	cookies["expire_country"] = "氪星"
	expiration := time.Now()
	expiration = expiration.AddDate(0,0,3)
	bvalue := []byte(cookies["tempname"])
	encodeString := base64.StdEncoding.EncodeToString(bvalue)
	nameCookie := http.Cookie{Name:"tempname",Value:encodeString}  // 临时的Cookie


	bvalue = []byte(cookies["expire_country"])
	encodeString = base64.StdEncoding.EncodeToString(bvalue)
	countryCookie := http.Cookie{Name:"expire_country",Value:encodeString,Expires:expiration}


	http.SetCookie(w,&nameCookie)
	http.SetCookie(w,&countryCookie)

	fmt.Fprintf(w,"write cookie success")


}
func readTemporaryCNCookies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"<html>")
	for _,cookie:=range r.Cookies() {
		if cookie.Name != "tempname" && cookie.Name != "expire_country" {
			continue
		}
		fmt.Fprint(w,cookie.Name)
		fmt.Fprint(w,"=")
		decodeBytes,_ := base64.StdEncoding.DecodeString(cookie.Value)
		value := string(decodeBytes)
		fmt.Fprint(w,value)
		fmt.Fprint(w,"<br>")
	}
	fmt.Fprint(w,"</html>")

}
func main() {
	// 指定路由和回调函数
	http.HandleFunc("/writeTemporaryCNCookies",writeTemporaryCNCookies)
	http.HandleFunc("/readTemporaryCNCookies",readTemporaryCNCookies)
	fmt.Println("服务器已经启动")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}

