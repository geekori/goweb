package main

import (
	"net/http"
	"html/template"
	"fmt"
	"log"
	"os"
	"io"
)

/*

上传文件


上传文件经过如下3步：
1.  在Web页面选择一个文件，然后上传
2.  在服务端读取上传文件的数据（字节流）
3.  将文件数据写到服务端的某一个文件中




 */

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024*1024)  // 最多在内存中一次处理1MB的数据
	file, handler,err := r.FormFile("uploadfile")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()  //  延迟关闭文件（在uploadFile函数结束时关闭文件）

	fmt.Fprintf(w, "%v",handler.Header)

	// 打开服务端的文件

	f, err := os.OpenFile("./upload/" + handler.Filename, os.O_WRONLY | os.O_CREATE,0666)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	io.Copy(f, file)


}
func showUploadfilePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		t,_:=template.ParseFiles("./src/form/uploadfile.html")
		t.Execute(w,nil)
	}
}

func main() {
	// 指定路由和回调函数
	http.HandleFunc("/",showUploadfilePage)
	http.HandleFunc("/upload",uploadFile)
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}