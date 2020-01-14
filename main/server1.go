package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func userAgent(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.UserAgent()) // 将结果写入response
}

// 请求处理函数
func path(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL.Path[1:]) // 将结果写入response
}

func notFound(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}

	t, err := template.ParseFiles("main/404.html")
	if (err != nil) {
		log.Println(err)
	}
	t.Execute(w, "not found")
}

func main() {
	http.HandleFunc("/", path) // 添加路由处理函数
	http.HandleFunc("/user-agent", userAgent)
	http.HandleFunc("/404", notFound)
	http.ListenAndServe(":8080", nil) // 监听本机端口
	// http.ListenAndServe("", nil) // 默认为':80'
}
