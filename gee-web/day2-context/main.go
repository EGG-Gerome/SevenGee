package main

import (
	"github.com/EGG-Gerome/gee-web/day1-http-base/base3/gee"
	"net/http"
)

func indexHandler(c *gee.Context) {
	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
}

func main() {
	r := gee.New()
	r.GET("/", indexHandler)
}
