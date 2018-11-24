package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"strconv"
)

/*

读写多个Cookie


 */

func writeCookies(w http.ResponseWriter, r *http.Request)  {
    var cookies map[string]interface{}
    cookies = make(map[string]interface{})
    cookies["name"] = "Bill"
    cookies["age"] = 36
    cookies["salary"] = 2000
    cookies["country"] = "China"
	cookies["name1"] = "小明"

	expiration := time.Now()
	expiration = expiration.AddDate(0,0,3)

	for key,value := range cookies {
		v,yes := value.(string)

		//  转换失败
		if !yes {
			var intV = value.(int)
			v = strconv.Itoa(intV)
		}
		cookie := http.Cookie{Name:key,Value:v,Expires:expiration}
		http.SetCookie(w,&cookie)

	}

	fmt.Fprintf(w,"write cookie success")


}
func readCookies(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,"<html>")

    for _,cookie := range r.Cookies() {
    	fmt.Fprint(w,cookie.Name)
    	fmt.Fprint(w,"=")
    	fmt.Fprint(w,cookie.Value)
    	fmt.Fprint(w,"<br>")

	}
	fmt.Fprint(w,"</html>")

}
func main() {
	// 指定路由和回调函数
	http.HandleFunc("/writeCookies",writeCookies)
	http.HandleFunc("/readCookies",readCookies)
	fmt.Println("服务器已经启动")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}


