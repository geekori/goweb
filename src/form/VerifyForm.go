package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
	"strconv"
	"regexp"
)

/*

校验表单数据

1. 必填字段和数字字段
2. 汉语   \p{Han}   \p{Greek}
3. 英文和数字
4. EMail
5. 日期
6. 身份证号
7. 单选列表
8. 单选按钮（radio）
9. 多选按钮（checkbox）

数字校验：（1） 校验输入的内容是否为数字   （2）校验数字的范围

1. 必填字段
2. 数字
3. 汉语
4. 英文和数字
5. Email
6. 手机号
7. 身份证号
8. combobox
9. checkbox
10.radio

 */

func verify(w http.ResponseWriter, r *http.Request) {
	//  保证只能处理POST请求
	if r.Method == "GET" {
		return
	}
	r.ParseForm()   //  分析客户端的body数据
	fmt.Println(r.Form)
    //  校验必填字段
	names,ok := r.Form["name"]
	if ok == false {
		return
	}
	name := names[0]
	if len(name) == 0 {
		fmt.Fprintf(w, "姓名不能为空")
		return
	}

    // 校验数字字段

    //  将字符串转换为数字
    intValue,err := strconv.Atoi(r.Form.Get("number"))
    if err != nil {
    	fmt.Fprintf(w,"输入的不是整数")
    	return
	}

	// 要求输入的整数不能大于100
	if intValue > 100 {
		fmt.Fprintf(w,"输入的数字不能大于100")
		return
	}

	//  用正则表达式校验汉语
	chinese := r.Form.Get("chinese")
	if m,_ := regexp.MatchString("^\\p{Han}+$",chinese); !m {
		fmt.Fprintf(w,"必须输入汉语")
		return
	}

	//  用正则表达式校验英文和数字的混合输入
	engnum := r.Form.Get("engnum")
	if m,_ := regexp.MatchString("^[a-zA-Z0-9]+$",engnum); !m {
		fmt.Fprintf(w,"必须输入英文或数字")
		return
	}

	// 用正则表达式校验EMail
	email := r.Form.Get("email")
	//  bill@google.com
	if m,_ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`,email); !m {
		fmt.Fprintf(w,"Email格式不正确")
		return
	}

    //  用正则表达式校验日期字段

    dateField := r.Form.Get("datefield")
    //  日期： 2001-1-1   2001-01-12   2019-3-5
	if m,_ := regexp.MatchString(`^\d{4}-\d{1,2}-\d{1,2}$`,dateField); !m {
		fmt.Fprintf(w,"日期格式不正确")
		return
	}

	//  用正则表达式校验身份证号
	idcard := r.Form.Get("idcard")
	//  校验18位身份证号
	if m,_ := regexp.MatchString(`^(\d{17})([0-9]|X)$`,idcard); !m {
		// 校验15位身份证号
		if m,_ := regexp.MatchString(`^\d{15}$`,idcard); !m {
			fmt.Fprintf(w,"请输入15位或18位身份证号")
			return
		}

	}
    // 校验单选列表
    fruit := r.Form.Get("fruit")
    fruits := []string{"peach","banana","apple"}
    errorFlag := true
    for _,item := range fruits {
    	if item == fruit {
    		errorFlag = false
		}
	}
	if errorFlag {
		fmt.Fprintf(w,"没有这种水果")
		return
	}
    //  校验单选按钮
    sexValue := r.Form.Get("sex")
    sexValues := []string{"1","2"}

    for _,v := range sexValues {
    	if v == sexValue {
    		errorFlag = false
		}
	}
    if errorFlag {
    	fmt.Fprintf(w, "你确定除了男女，还有其他性别吗？")
    	return
	}
    errorFlag = false
	//  校验多选按钮
    var valueMap = map[string]int{"football":1,"basketball":1,"tennis":1}
    interestValues := r.Form["interest"]
    for _,v := range interestValues {
    	_,ok:= valueMap[v]
    	if !ok {
    		errorFlag = true
		}
	}
    if errorFlag {
    	fmt.Fprintf(w,"兴趣爱好不存在")
    	return
	}
	fmt.Fprintf(w,"success")


}
func showForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t,_:=template.ParseFiles("/MyStudio/video_src/goweb/src/form/form.html")
		t.Execute(w,nil)
	}
}

func main() {
	// 指定路由和回调函数
	http.HandleFunc("/",showForm)
	http.HandleFunc("/verify",verify)
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900")
	//  启动HTTP服务，并监听端口号
	err := http.ListenAndServe(":8900",nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}