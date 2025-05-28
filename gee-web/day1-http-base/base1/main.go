package base1

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler) // HandleFunc 用于注册路由和处理函数的绑定关系。
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":1103", nil))
	// nil参数表示使用默认的路由多路复用器（DefaultServeMux）
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	// w http.ResponseWriter 是用于向客户端发送 HTTP 响应的接口
	// req *http.Request：客户端发送的 HTTP 请求对象，包含请求行、头部、Body 等信息
	fmt.Fprintf(w, "URL.PATH = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

// test cmd
// $ curl http://localhost:1103/
// URL.Path = "/"
// $ curl http://localhost:1103/hello
// Header["Accept"] = ["*/*"]
// Header["User-Agent"] = ["curl/7.54.0"]
