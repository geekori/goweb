package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

/*

读取对象JSON文档

 */

type MyProduct1 struct {
	ProductName string
	Price float64
}

type MyProducts1 struct {
	Products []MyProduct1
	Count int
	Country string
}


func main() {
	file,err := os.Open("./src/documents/obj.json")
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
	var p MyProducts1
	json.Unmarshal(data,&p)

	fmt.Println(p.Count)
	fmt.Println(p.Country)

	for i :=0;i < len(p.Products);i++ {
		fmt.Println("产品名称：" + p.Products[i].ProductName)
		fmt.Printf("产品价格：%.2f\n",p.Products[i].Price)
	}
}


