package main

import (
	"strconv"
	"math/rand"
	"time"
	"fmt"
	"net/http"
	"log"
)

/*
在服务端使用Session

1. 服务端根据SessionID检测客户端是否第一次访问服务端

2. 如果服务端没有检测到SessionID，可能有下面的两种情况
   （1）Session到期，服务端已经删除了该Session
    (2)SessionID是假的

3.  如果服务端找到了SessionID，那么客户端就不是第一次访问，然后根据业务逻辑进行下一步处理

4.  如果客户端没有发送SessionID，或者SessionID没找到。服务端就会产生新的SessionID

5.  将产生的SessionID通过HTTP响应头发送给客户端


如何产生SessionID

0 - 9  a - z A - Z
 */
// 产生指定长度的随机字符串

func generateRandomStr(n int) string {
	var strarr = [62]string{}
	result := ""
	//  填充0到9
	for i:=0; i< 10;i++ {
		strarr[i] = strconv.Itoa(i)
	}

	//  填充a到z
	for i:=0;i<26;i++ {
		strarr[i + 10] = string(i+97)
	}

	//  填充A到Z
	for i:=0;i<26;i++ {
		strarr[36+i] = string(i + 65)
	}
    source := rand.NewSource(time.Now().Unix())
    r:=rand.New(source)
    for i:=0;i<n;i++ {
    	index := r.Intn(len(strarr))
    	result = result + strarr[index]
	}


	return result
}

var sessions map[string]string

func writeSessionID(w http.ResponseWriter) {
	sessionId := generateRandomStr(20)
	sessions[sessionId] = generateRandomStr(100)
	expiration := time.Now()
	expiration = expiration.AddDate(0,0,40)
	cookie := http.Cookie{Name:"mysession_id",Value:sessionId,Expires:expiration}
	http.SetCookie(w,&cookie)

}
func mysession(w http.ResponseWriter, r *http.Request) {
	cookie,err := r.Cookie("mysession_id")
	// 该用户已经提交了SessionID
	if err == nil {
		//  校验SessionID
		data,exists := sessions[cookie.Value]
		//  服务端存在这个SessionID，说明用户不是第一次访问服务端
		if exists {
			fmt.Fprint(w,"该用户已经访问过服务器了：" + data)
		} else {
			// 需要重新产生SessionID
			writeSessionID(w)
			fmt.Fprint(w,"该用户第一次访问服务器")
		}
	} else {
		// 需要重新产生SessionID
		writeSessionID(w)
		fmt.Fprint(w,"该用户第一次访问服务器")
	}

}
func main() {
	sessions = make(map[string]string)
	// 指定路由和回调函数
	http.HandleFunc("/",mysession)

	fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}

