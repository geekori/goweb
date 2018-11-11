package main

import (
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io"
	"net/http"
	"io/ioutil"
)

/*
模拟浏览器上传文件
 */

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile",filename)
	if err != nil {
		fmt.Println("文件写入缓存错误")
		return err
	}

	//  打开待上传的文件
	fh,err := os.Open("./src/form/" + filename)
	if err != nil {
		fmt.Println("打开文件错误")
		return err
	}

	defer  fh.Close()

	_,err = io.Copy(fileWriter, fh)  //  将硬盘上的文件读到客户端机器上的内存中
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	fmt.Println(contentType)

	bodyWriter.Close()

	response, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	//  读取服务端返回的数据
	resp_body,err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(response.Status)
	fmt.Println(string(resp_body))
	return nil



}

func main() {
	targetUrl := "http://localhost:8900/upload"
	filename := "uploadclient.go"
	postFile(filename, targetUrl)
}