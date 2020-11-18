package main

import (
	"fmt"
	"net/http"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "----")
}

type var1 struct {
}

func (v *var1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `var1`)
}

type var2 struct {
}

func (v *var2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `var2`)
}

// 处理器函数
func var3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "var3")
}

func main() {
	// handler := handler{}
	var1 := var1{}
	var2 := var2{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// 将自定义处理器和服务器绑定 默认多路复用器被替代 所有请求由自定义处理器处理
		// Handler: &handler,
	}

	http.Handle("/var1", &var1)
	http.Handle("/var2", &var2)

	// cert.pem是SSL证书 key.pem是服务器的私钥
	// server.ListenAndServeTLS("cert.pem", "key.pem")
	server.ListenAndServe()
}
