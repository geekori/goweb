package main

import (
	"regexp"
	"fmt"
)

/*

正则表达式的查找函数

Find
FindAll
FindIndex
...


 */
func main() {
    a:="Space, the final frontier. fronxier These are the voyages of the starship Enterprise. Its five-year mission: To explore strange new worlds. To seek out new life and new civilizations. To boldly go where no man has gone before."

    re,_:=regexp.Compile("[a-z]{2,4}")  // ab  abc

    // 查找符合正则的第一个子字符串
    one := re.Find([]byte(a))
    fmt.Println("Find:" ,string(one))   // pace

    // FindAll
    // n:-1 表示查找所有符合条件的子字符串    n: >0 （m)   表示查找前m个数符合条件的子字符串
    all := re.FindAll([]byte(a),3)
    fmt.Println("FindAll",all)
    fmt.Println("FindAll:", string(all[0]))
	fmt.Println("FindAll:", string(all[1]))
	fmt.Println("FindAll:", string(all[2]))

    // FindIndex
    // 查找符合条件的子字符串的Index位置，包括起始位置和结束位置

    index := re.FindIndex([]byte(a))
    fmt.Println("FindIndex",index)  // pace的起始位置和结束位置

    // FindAllIndex

	allIndex := re.FindAllIndex([]byte(a),-1)
	fmt.Println("FindAllIndex",allIndex)


	//  FindSubmatch
	//  返回数组：第1个元素表示整个匹配的字符串，第2个元素表示第1个()内匹配的值，第3个元素表示第2个()内匹配的值
	// frontier
	re2,_:=regexp.Compile("fr(o)(n.)ier")

	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println(submatch)
	for _,v := range submatch {
		fmt.Println(string(v))
	}

	// FindAllSubmatch
	submatchall := re2.FindAllSubmatch([]byte(a),-1)
	fmt.Println(submatchall)
	for _,v1 := range submatchall {
		for _,v2 := range v1 {
			fmt.Println(string(v2))
		}
	}

	// FindSubmatchIndex([]byte(a))
	// FindAllSubmatchIndex([]byte(a),-1)
}

