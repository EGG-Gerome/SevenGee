package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
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

//	func askName(c *gin.Context) {
//		c.String(http.StatusOK, "Who are you?")
//	}
//
//	func sayHelloToName(c *gin.Context) {
//		name := c.Param("name")
//		c.String(http.StatusOK, "Hello %s", name)
//	}
//
//	func queryNameAndRole(c *gin.Context) {
//		name := c.Query("name") // 查不到的默认值为空
//		role := c.DefaultQuery("role", "teacher")
//		c.String(http.StatusOK, "%s is a %s", name, role)
//	}
//
//	func postUsernameAndPassword(c *gin.Context) {
//		// 1. 从表单中获取用户名和密码
//		username := c.PostForm("username")
//		password := c.DefaultPostForm("password", "123456") // 可以设置默认值，但密码不要设置初始值
//		// 2. 验证输入是否合法
//		if username == "" || password == "" {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": "username and password can't be empty",
//			})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{
//			"username": username,
//			"password": password,
//		})
//	}
//
//	func myMap(c *gin.Context) {
//		// 1. 获取URL查询参数中的"ids"映射（适用于GET请求）
//		ids := c.QueryMap("ids")
//
//		// 2. 获取表单POST参数中的"names"映射（适用于POST请求）
//		names := c.PostFormMap("names")
//
//		// 3. 将两个映射作为JSON响应返回
//		c.JSON(http.StatusOK, gin.H{
//			"ids":   ids,
//			"names": names,
//		})
//	}
//
//	func redirect(c *gin.Context) {
//		c.Redirect(http.StatusMovedPermanently, "/index")
//	}
//
//	func queryRelativePath(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"path": c.FullPath(),
//		})
//	} // group routes
//func updateSingleFile(c *gin.Context) {
//	file, err := c.FormFile("file")
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	// c.SaveUploadedFile(file, dst)
//	c.String(http.StatusOK, "%s uploaded!", file.Filename)
//}
//func updateMultipartFile(c *gin.Context) {
//	// Multipart form
//	form, err := c.MultipartForm()
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": "解析表单失败: " + err.Error(),
//		})
//		return
//	}
//	// 获取文件列表
//	files := form.File["update[]"]
//	for _, file := range files {
//		log.Println(file.Filename)
//		// c.Save
//	}
//}

//func main() {
//gin.SetMode(gin.ReleaseMode) // 切换至发布模式
//r := gin.Default()
//
//r.GET("/", askName)
//r.GET("/user/:name/*role", sayHelloToName) // 无论/:name/后面输入了什么，都会调用sayHello 匹配user?name=xxx&role=xxx, role可选
//r.GET("/users", queryNameAndRole)
//r.POST("/form", postUsernameAndPassword)
//r.POST("/map", myMap)
//r.GET("/redirect", redirect)
//r.GET("/goIndex", func(c *gin.Context) {
//	c.Request.URL.Path = "/"
//	r.HandleContext(c)
//})
//
//defaultHandler := queryRelativePath	// group routes
//// group: v1
//v1 := r.Group("/v1")
//{
//	v1.GET("/posts", defaultHandler)
//	v1.GET("/series", defaultHandler)
//}
//v2 := r.Group("/v2")
//{
//	v2.GET("/posts", defaultHandler)
//	v2.GET("/series", defaultHandler)
//}
//r.POST("/upload1", updateSingleFile)
//r.POST("/upload2", updateMultipartFile)
//r.Run(":1103")
//}

// HTML Template
//func main() {
//	r := gin.Default()
//
//	type student struct {
//		Name string
//		Age  int8
//	}
//
//	r.LoadHTMLGlob("templates/*.tmpl")
//
//	stu1 := &student{Name: "Emily", Age: 20}
//	stu2 := &student{Name: "Gerome", Age: 20}
//	r.GET("/arr", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
//			"title":  "Gin",
//			"stuArr": [2]*student{stu1, stu2},
//		})
//	})
//	r.Run(":1103") // 启动一个 HTTP 服务器，并在指定端口上持续监听请求。这会阻塞主线程，防止程序退出。
//}

// 定义中间件函数：logger
// 返回值类型为 gin.HandlerFunc（Gin中间件的标准类型）
func logger() gin.HandlerFunc {
	return func(c *gin.Context) { // 中间件的具体逻辑
		t := time.Now() // 记录请求开始时间

		// 给请求上下文（Context）设置一个键值对（键为"Egg"，值为"littleEgg"）
		// 后续的处理函数或中间件可以通过 c.Get("Egg") 获取该值
		c.Set("Egg", "littleEgg")

		// ---------------------- 请求处理前 ----------------------
		c.Next() // 调用Next()表示继续处理后续的中间件或路由处理函数
		// 如果不调用Next()，请求会在此处被拦截，不会继续往后执行

		// ---------------------- 请求处理后 ----------------------
		latency := time.Since(t) // 计算请求处理耗时
		log.Print(latency)       // 打印耗时（通常用于日志记录）
	}
}

// ---------------------- 以下是示例函数（需自行实现） ----------------------
// 自定义中间件（示例）
func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 中间件逻辑（例如：性能基准测试）
		c.Next()
	}
}

// 路由处理函数（示例）
func benchEndpoint(c *gin.Context) {
	c.String(200, "Benchmark endpoint")
}

// 权限验证中间件（示例）
func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证用户权限（例如：检查JWT令牌）
		c.Next()
	}
}

// 登录处理函数（示例）
func loginEndpoint(c *gin.Context) {
	c.String(200, "Login endpoint")
}

// 提交处理函数（示例）
func submitEndpoint(c *gin.Context) {
	c.String(200, "Submit endpoint")
}
func main() {
	r := gin.Default() // 创建Gin引擎实例（默认包含Logger和Recovery中间件）

	// ---------------------- 全局中间件 ----------------------
	// 使用内置的Logger中间件（记录请求日志）
	r.Use(gin.Logger())
	// 使用内置的Recovery中间件（捕获panic并恢复程序，防止服务器崩溃）
	r.Use(gin.Recovery())

	// ---------------------- 单个路由中间件 ----------------------
	// 给"/benchmark"路由单独添加自定义中间件MyBenchLogger()
	// 中间件顺序：先执行MyBenchLogger()，再执行路由处理函数benchEndpoint
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// ---------------------- 路由分组中间件 ----------------------
	// 创建名为authorized的路由分组（路径前缀为"/"，即与根路径同级）
	authorized := r.Group("/")
	// 给该分组添加自定义中间件Authorized()（例如：权限验证）
	authorized.Use(Authorized())
	{ // 分组内的路由会继承中间件Authorized()
		authorized.POST("/login", loginEndpoint)   // 登录接口
		authorized.POST("/submit", submitEndpoint) // 提交接口
	}

	// ---------------------- 启动服务器 ----------------------
	// 监听端口1103（默认IP为0.0.0.0）
	// 等价于 r.Run(":1103")
	r.Run(":1103")
}
