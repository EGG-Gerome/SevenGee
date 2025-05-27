package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func main() {
//	// 创建一个默认的Gin引擎实例，自动包含两个关键中间件：
//	//Logger：记录 HTTP 请求日志
//	//Recovery：捕获 panic 并恢复，防止服务器崩溃
//	r := gin.Default() // gin.Default()会返回一个*gin.Engine类型的实例，该实例代表着整个 Web 服务器。
//
//	// 定义一个GET请求的路由处理函数，当访问根路径"/"时执行
//	r.GET("/", func(c *gin.Context) { // 处理函数（匿名函数）接收一个*gin.Context类型的参数，它封装了 HTTP 请求和响应的上下文
//		// 参数c *gin.Context是在路由处理函数里传递进来的，它封装了单个 HTTP 请求的上下文信息，像请求头、请求参数以及响应等内容都包含在内
//
//		// 返回HTTP 200状态码和字符串响应
//		c.String(200, "Bello, Egg!")
//	})
//
//	// 启动服务器，监听并服务于0.0.0.0:8080 地址
//	r.Run() // 可以通过参数指定其他地址，例如：r.Run(":8081") 或 r.Run("localhost:8080")
//}

func askName(c *gin.Context) {
	c.String(http.StatusOK, "Who are you?")
}

func main() {
	r := gin.Default()
	r.GET("/", askName)
	r.Run()
}
