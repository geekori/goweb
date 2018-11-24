package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"encoding/base64"
)

/*

读写中文Cookie


 */

func writeCNCookies(w http.ResponseWriter, r *http.Request)  {

    var cookies map[string]string
    cookies = make(map[string]string)
    cookies["name"] = "王军"
    cookies["country"] = "中国"
	expiration := time.Now()
	expiration = expiration.AddDate(0,0,3)

	for key,value:=range cookies {
		bvalue := []byte(value)
		encodeString := base64.StdEncoding.EncodeToString(bvalue)
		cookie := http.Cookie{Name:key,Value:encodeString,Expires:expiration}
		http.SetCookie(w,&cookie)
	}

	fmt.Fprintf(w,"write cookie success")


}
func readCNCookies(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,"<html>")
    for _,cookie:=range r.Cookies() {
    	if cookie.Name != "name" && cookie.Name != "country" {
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
	http.HandleFunc("/writeCNCookies",writeCNCookies)
	http.HandleFunc("/readCNCookies",readCNCookies)
	fmt.Println("服务器已经启动")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}


