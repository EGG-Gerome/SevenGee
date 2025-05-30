package main

import (
	"fmt"
	"net/http"

	"github.com/EGG-Gerome/gee-web/day1-http-base/base3/gee"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
func main() {
	r := gee.New()
	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)
	r.RUN(":1103")
}
