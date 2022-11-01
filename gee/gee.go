package gee

import (
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	// 实现ServeHTTP的接口
	router *router
}
// 创建一个新的Engine
func New() *Engine {
	return &Engine{
		newRouter(),
	}
}

func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc){
	engine.router.addRoute(method, pattern, handlerFunc)
}
// 一共有6种http请求的方法
// GET POST PUT HEAD DELETE CONNECT OPTIONS TRACE PATCH
// 添加了get请求的方法
func (engine *Engine) Get(pattern string, handler HandlerFunc){
	engine.addRoute("GET", pattern, handler)
}

// 添加了POST请求的方法
func (engine *Engine) Post(pattern string, handler HandlerFunc){
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) RUN(addr string) (err error){
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
	c := newContext(w, req)
	engine.router.handle(c)
}



