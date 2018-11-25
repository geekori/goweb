package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

/*

读取数组JSON文档


JSON：数组和对象

数组:[...]
对象:{...}

如果不使用tag，必须保证struct中的属性名和json文档对应的属性名相同（大小写不敏感）

如果名字不一致，需要使用tag进行映射。tag本身也是大小写不敏感的

 */
// struct tag
type MyProduct struct {
	ProductName string  `json:"productName1"`
	Price float64
}

type MyProducts []MyProduct



func main() {
	file,err := os.Open("./src/documents/array.json")
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
	var p MyProducts
	json.Unmarshal(data, &p)
	for i:=0;i<len(p);i++{
		fmt.Println("产品名称：" + p[i].ProductName)
		fmt.Printf("产品价格：%.2f\n",p[i].Price)
	}
}


