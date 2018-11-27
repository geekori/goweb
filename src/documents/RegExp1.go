package main

import (
	"os"
	"fmt"
	"regexp"
)

/*
用正则表达式匹配文本

数字：^[0-9]+$
IP：xxx.xxx.xxx.xxx   ^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$
Email：abc@126.com    ^[a-zA-Z0-9-_]+@[a-zA-Z0-9-_]+.(com|com.cn|net|org|cn)$

regexp   MatchString(正则表达式,待匹配的字符串）如果匹配成功，返回true，否则返回false
 */


func main() {
    if len(os.Args) == 1 {
    	fmt.Println("Usage:regexp1 [string]")
    	os.Exit(1)
	} else if m,_:=regexp.MatchString("^[0-9]+$",os.Args[1]);m {
		fmt.Println("是数字")

	} else if m,_:=regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$",os.Args[1]);m{
		fmt.Println("是IP")
	} else if m,_ := regexp.MatchString("^[a-zA-Z0-9-_]+@[a-zA-Z0-9-_]+.(com|com.cn|net|org|cn)$",os.Args[1]);m {
		fmt.Println("是Email")
	} else {
		fmt.Println("是普通字符串")
	}

}