package main

import (
	"net/http"
	"fmt"
	"log"
	"session_library"
	"time"
	"html/template"
)


func count(w http.ResponseWriter, r *http.Request) {

	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		ct = 1
	} else {
		ct = ct.(int) + 1
	}
	sess.Set("countnum", ct)
	t, _ := template.ParseFiles("./src/session/count.html")

	t.Execute(w, ct)

	w.Header().Set("Content-Type", "text/html")
}
var globalSessions *geekori_session.Manager
func init() {
	// redis
	globalSessions, _ = geekori_session.NewManager("memory", "gosessionid", 3600)
	// goroutine，检测Session是否到期
	go globalSessions.GC()
}
func main() {

	http.HandleFunc("/count", count) //设置访问的路由
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入http://localhost:8900/count")
	err := http.ListenAndServe(":8900", nil) //设置监听的端口
	fmt.Println("监听之后")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}