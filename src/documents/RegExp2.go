package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

/*

用正则表达式替换文本内容

regexp.Compile

re.ReplaceAllStringFunc(src,target)

 */
func main() {
	resp,err := http.Get("http://www.jd.com")
	if err != nil {
		fmt.Println("http get error")
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	// 将HTML标签转换为小写
	// \s：匹配任何的空白符（空格、制表符、换页符等）
	// \S：匹配任何非空白符，
	re,_ := regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllStringFunc(src,strings.ToUpper)


	//  去掉所有的style标签
	re,_ = regexp.Compile(`<STYLE[\S\s]+?</STYLE>`)
	src = re.ReplaceAllString(src,"")

	//  去掉所有的script标签
	re,_ = regexp.Compile(`<SCRIPT[\S\s]+?</SCRIPT>`)
	src = re.ReplaceAllString(src,"")

	//  去除所有尖括号内的HTML代码，并替换成换行符
	re,_ = regexp.Compile(`<[\s\S]+?>`)
	src = re.ReplaceAllString(src,"\n")

	//  去掉连续的换行符
	re,_ = regexp.Compile(`\s{2,}`)
	src = re.ReplaceAllString(src,"\n")
	fmt.Println(strings.TrimSpace(src))
}
