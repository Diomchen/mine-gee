package gee

import (
	"log"
	"net/http"
)

// 为什么这里需要用小写的router，而不是Router？
// 因为Router是Go标准库里的一个包名，所以为了避免冲突，这里用小写的router。
type router struct {
	handlers map[string]HandlerFunc
}

// 为什么同包内可以直接调用的newRouter()方法，这里却需要用小写的newrouter()方法？
// 因为Go里的大写字母开头的函数名一般都是包外可见的，而小写字母开头的函数名一般都是包内可见的。
// 所以为了避免冲突，这里用小写的newrouter()方法。
func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("adding route %s %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) hendle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 page not found")
	}
}
