package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}
// 对Web服务来说，无非是根据请求*http.Request，构造响应http.ResponseWriter
type Context struct {
	// 原始的需求
	Writer 	http.ResponseWriter
	Req 	*http.Request
	// request 的参数
	Path	string
	Method	string
	Params map[string]string 	// 查找一部分的路由
	// 状态码
	StatusCode	int
	// middleware
	handlers []HandlerFunc
	index int	// 当前执行到第几个中间件
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
		index: -1,
	}
}
// 中间件到next函数 执行下一个中间件
// 中间件使用了next先执行后面的handle的时候，采用的就是循环外的++
// 如果没用next的话。就是用的循环内的++
func (c *Context) Next(){
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// 可以获取？之后的参数
func (c *Context) PostForm(key string) string{
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string{
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}
func (c *Context) String(code int, format string, values ...interface{}){
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil{
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte){
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}