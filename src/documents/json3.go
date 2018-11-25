package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

/*

将JSON文档映射到interface{}上

 */




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
	var p interface{}
	json.Unmarshal(data, &p)
	fmt.Println(p)
    pp := p.(map[string]interface{})
    fmt.Println(pp["count"])
    fmt.Println(pp["country"])
	fmt.Println(pp["products"])

    products := pp["products"].([]interface{})

    for i:=0; i < len(products);i++ {
    	product := products[i].(map[string]interface{})
    	fmt.Println(product["productName"])
    	fmt.Printf("%.2f\n",product["price"])
	}


}


