package gee

import (
	"fmt"
	"net/http"
)

// day1-http-base

/*
定义请求处理函数类型：
作用：定义一个函数类型 HandlerFunc，他接受两个参数

	http.RespondWriter：用于向客户端发送请求
	*http.Request： 客户端的HTTP 请求（指针类型）

意义：所有的路由处理函数都必须符合这个签名，确保接口统一
*/
// 用来规定传入的参数（这里的参数是函数）是这个模式的（函数会有这两个传入）
type HandlerFunc func(http.ResponseWriter, *http.Request)

/*
定义路由器核心结构 Engine
router：一个映射表(map)，键为 URL 路径（如"/hello"），值为对应的处理函数(HandlerFunc 类型)
作用：Engine 是路由器的核心，负责存储路径与处理函数的映射关系，并在接收到请求时根据路径找到相应的处理函数
*/
type Engine struct {
	router map[string]HandlerFunc
}

/*
构造函数 New()
作用：创建并初始化一个 Engine 实例，返回其指针
关键点：

	make(map[string]HandlerFunc)：初始化 map，避免空指针异常
	返回指针 *Engine 确保所有调用者共享同一个路由器实例（单例模式）
*/
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

/*
method（HTTP 请求方法）

	类型：string (eg "GET","POST","PUT")
	作用：指定请求的 HTTP 方法，区分对同一 URL 的不同操作

pattern（URL路径模式）

	类型：string(eg "/","/hello","/user:id")
	作用：指定请求的 URL 路径或路径模板，支持静态路径和动态参数
*/
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// RUN defines the method to start a http server
func (engine *Engine) RUN(addr string) (err error) {
	return http.ListenAndServe(addr, engine) // 第二个参数必须是一个实现了 http.Handler 接口的对象。
}

// 第二个参数必须是一个实现了 http.Handler 接口的对象。
// http.Handler 接口定义：
//
//	type Handler interface {
//		ServeHTTP(w ResponseWriter, r *Request)
//	}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOTFOUND: %s\n", req.URL)
	}
}

// day2-context
type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code        // 记录状态码
	c.Writer.WriteHeader(code) // 发送状态码
}
