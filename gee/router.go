package gee

import "net/http"

// 构造关于路由的方法 方便对router功能进行增强

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router{
	return &router{make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handlerFunc HandlerFunc){
	key := method + "-" + pattern
	r.handlers[key] = handlerFunc
}

func (r *router) handle(c *Context){
	key := c.Method +"-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c.Writer, c.Req)
	}else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}