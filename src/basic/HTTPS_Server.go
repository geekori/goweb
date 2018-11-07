package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

/*
编写HTTPS服务器

HTTPS = HTTP + Secure（安全）

RSA进行加密

SHA进行验证

秘钥和证书

生成秘钥文件
openssl genrsa -out server.key  2048

生成证书文件
openssl req -new -x509 -sha256 -key server.key -out server.crt  -days 3650

下载Windows版本的OpenSSL工具

https://slproweb.com/products/Win32OpenSSL.html


1.  域名型SSL证书（DV SSL）：自动签发

2.  企业型SSL证书（OV SSL）：只需要提供合法手续就能签发

3. 增强型SSL证书（EV SSL）：银行或金融机构


Https





 */

 func httpsEcho(w http.ResponseWriter, r *http.Request) {
	 fmt.Println("Path:",r.URL.Path)
	 fmt.Println("Url:",r.URL)   // /a/b/
	 fmt.Println("Host",r.Host)  // localhost:8900
	 fmt.Println("Header:", r.Header)
	 fmt.Println("Method:",r.Method)
	 fmt.Println("Proto：", r.Proto)

	 fmt.Println("UserAgent:", r.UserAgent())

	 scheme := "http://"
	 if r.TLS != nil {
		 scheme = "https://"
	 }
	 fmt.Println("完整的请求路径：" + strings.Join([]string{scheme,r.Host,r.RequestURI},""))
	 fmt.Fprintf(w,"Hello Go Web")
 }

 func main() {
 	http.HandleFunc("/",httpsEcho)
 	fmt.Println("HTTPS服务器已经启动")
 	err := http.ListenAndServeTLS(":4321","/MyStudio/video_src/goweb/src/basic/server.crt","/MyStudio/video_src/goweb/src/basic/server.key",nil)
	 if err != nil {
		 log.Fatal("ListenAndServe:",err)
	 }
 }