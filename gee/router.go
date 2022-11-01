package gee

import (
	"net/http"
	"strings"
)

// 构造关于路由的方法 方便对router功能进行增强

type router struct {
	roots		map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router{
	return &router{
		make(map[string]*node),
		make(map[string]HandlerFunc),
	}
}
// 拆分路由。构造path
func parsePattern(pattern string) []string{
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != ""{
			parts = append(parts, item)
			if item[0] == '*'{
				break
			}
		}
	}
	return parts
}


func (r *router) addRoute(method string, pattern string, handlerFunc HandlerFunc){
	key := method + "-" + pattern
	parts := parsePattern(pattern)

	_, ok := r.roots[method]
	if !ok {
		// 构建方法树
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern,parts,0)

	r.handlers[key] = handlerFunc
}
// 查找路由
func (r *router) getRoute(method string, path string) (*node, map[string]string){
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil,nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		//例如/p/go/doc匹配到/p/:lang/doc，解析结果为：{lang: "go"}
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			//static/css/geektutu.css匹配到/static/*filepath
			//解析结果为{filepath: "css/geektutu.css"}
			if part[0] == '*' {
				params[part[1:]] = strings.Join(searchParts[index:],"/")
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context){

	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		// 如果是c.path那么这时候找到的就是/hello/tyc
		// 如果是n.pattern 那么这时候找到的就是/hello/
		key := c.Method +"-" + n.pattern
		r.handlers[key](c)
	}else{
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
	//if handler, ok := r.handlers[key]; ok {
	//	handler(c.Writer, c.Req)
	//}else {
	//	c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	//}
}