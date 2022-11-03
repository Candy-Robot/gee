package gee

import (
	"fmt"
	"testing"
)

//func newTestRouter() *router {
//	r := newRouter()
//	r.addRoute("GET", "/", nil)
//	r.addRoute("GET", "/hello/:name", nil)
//	//r.addRoute("GET", "/hello/b/c", nil)
//	//r.addRoute("GET", "/hi/:name", nil)
//	//r.addRoute("GET", "/assets/*filepath", nil)
//	return r
//}
////func TestParsePattern(t *testing.T) {
////	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
////	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
////	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
////	if !ok {
////		t.Fatal("test parsePattern failed")
////	}
////}
//
//func TestGetRoute(t *testing.T) {
//	r := newTestRouter()
//	n, ps := r.getRoute("GET", "/hello/222")
//
//	if n == nil {
//		t.Fatal("nil shouldn't be returned")
//	}
//	if n.pattern != "/hello/:name" {
//		t.Fatal("should match /hello/:name")
//	}
//
//	if ps["name"] != "222" {
//		t.Fatal("name should be equal to 'geektutu'")
//	}
//
//	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
//
//}
/// 测试传入参数之后，是否是整个属性都会变化。得出结论是是的
func A(c *Tyc){
	fmt.Println("test A start")
	c.add()
	fmt.Println("test A end")
}
func B(c *Tyc){
	fmt.Println("test B start")
	//c.add()
	fmt.Println("test B end")
}
func C(c *Tyc){
	fmt.Println("这是之后需要执行的handle")
}
type yc func(tyc *Tyc)

type Tyc struct {
	index int
	testlist []yc
}

func (tyc *Tyc) add(){
	tyc.index++
	fmt.Printf("add 循环外+1 %d\n",tyc.index)
	for tyc.index < len(tyc.testlist){
		tyc.testlist[tyc.index](tyc)
		tyc.index++
		fmt.Printf("add 循环内 +1 index = %d\n", tyc.index)
	}
}

func TestAdd(t *testing.T){
	var tlist []yc = []yc{A,B, C}
	tyctest := &Tyc{
		index: -1,
		testlist: tlist,
	}

	tyctest.add()
	fmt.Printf("pass")
}

