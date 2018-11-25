/*

处理XML文档

XML、JSON、模板文件、一般的文本或二进制文件


读取XML文件

xml.Unmarshal
 */
package main

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
)

// struct tag
type Products struct {
	 // 根节点名称
     XMLName xml.Name  	`xml:"products"`
	 // 属性
	 MyCount int 		`xml:"count,attr"`
	 MyCountry string 	`xml:"country,attr"`
	 // 子节点
	 ProductArr []Product `xml:"product"`
}
type Product struct {
	XMLName xml.Name `xml:"product"`
	ProductName string `xml:"productName"`
	Price float64 `xml:"price"`
}
func main() {
	file,err := os.Open("./src/documents/products.xml")
	if err != nil {
		fmt.Printf("error:%v",err)
		return
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error:%v",err)
		return
	}
    products := Products{}
    err = xml.Unmarshal(data, &products)
    if err != nil {
    	fmt.Printf("error:%v", err)
    	return
	}

	fmt.Println(products.XMLName.Local)
    fmt.Println(products.MyCount)
    fmt.Println(products.MyCountry)
    for i:=0;i<len(products.ProductArr);i++{
    	fmt.Println(products.ProductArr[i].XMLName.Local)
    	fmt.Println(products.ProductArr[i].ProductName)
    	fmt.Println(products.ProductArr[i].Price)
	}


}