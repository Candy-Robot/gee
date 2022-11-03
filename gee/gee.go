package gee

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type RouterGroup struct{
	prefix		string			// 通过前缀来判断
	middlewares	[]HandlerFunc	// 添加中间件
	parent		*RouterGroup	// 可以有嵌套，要支持分组嵌套
	engine   	*Engine			// 每个组都有实现路由的能力
}
// 我们还可以进一步地抽象，将Engine作为最顶层的组
type Engine struct {
	*RouterGroup
	router *router		// 实现ServeHTTP的接口
	groups []*RouterGroup	// 存放所有的组
}

// 将中间件添加到组中
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// 创建一个新的Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 创建一个新的组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine		// 有实现路由的能力
	newGroup := &RouterGroup{	// 创建一个新的子组
		prefix: prefix,
		parent: group,			// 父组是之前的group
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)	// 将新的组加入到组集合中
	return newGroup
}

func (gourp *RouterGroup) addRoute(method string, comp string, handlerFunc HandlerFunc){
	pattern := gourp.prefix + comp		// 增加路由的时候需要将组的前缀加上
	// 由于Engine从某种意义上继承了RouterGroup的所有属性和方法，因为 (*Engine).engine 是指向自己的。
	gourp.engine.router.addRoute(method, pattern, handlerFunc)
}
// 一共有6种http请求的方法
// GET POST PUT HEAD DELETE CONNECT OPTIONS TRACE PATCH
// 添加了get请求的方法
func (group *RouterGroup) GET(pattern string, handler HandlerFunc){
	group.addRoute("GET", pattern, handler)
}

// 添加了POST请求的方法
func (group *RouterGroup) POST(pattern string, handler HandlerFunc){
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) RUN(addr string) (err error){
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
	// 通过前缀来判断是否属于某个组当中
	// 没有写组拼接的代码
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix){
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}



