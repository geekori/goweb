package main

import (
	"strings"
	"fmt"
)

/*
字符串的基本操作

1. 字符串包含
2. 字符串连接
3. 字符串分割
4. 字符串定位
5. 字符串重复n次
6. 去掉字符串的空格，并按空格将字符串分割成slice
7. 字符串首尾剪切
 */
func main() {
    // 1. 字符串包含
    fmt.Println(strings.Contains("hello world","wor"))
	fmt.Println(strings.Contains("hello world","wora"))
	fmt.Println(strings.Contains("hello world",""))
	fmt.Println(strings.Contains("",""))

    // 2. 字符串连接
    s := []string{"Bill","Mike","John","Mary"}
    fmt.Println(strings.Join(s," "))

    // 3. 字符串分割
    fmt.Printf("%q\n",strings.Split("a,b,c,d",","))
	fmt.Printf("%q\n",strings.Split("How are you"," "))
	fmt.Printf("%q\n",strings.Split("abcdefg",""))

    // 4. 字符串定位，返回子字符串的索引，如果未搜索到，返回-1
    fmt.Println(strings.Index("How are you?","ar"))
	fmt.Println(strings.Index("How are you?","ara"))

    // 5. 字符串重复n次
    fmt.Println("12" + strings.Repeat("0",10))
	fmt.Println(strings.Repeat("hello",10))

	// 6. 去掉字符串的空格，并按空格将字符串分割成slice
	fmt.Printf("Fields are:%q\n", strings.Fields("   How    are you ? "))


	// 7. 字符串首尾剪切，只要字符串首尾的字符包含在cutset中，就会被截取
	fmt.Printf("[%q]",strings.Trim("  !!* abcd      !!"," !*ac"))



}