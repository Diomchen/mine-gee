package gee

import (
	"net/http"
	"strings"
)

// 为什么这里需要用小写的router，而不是Router？
// 因为Router是Go标准库里的一个包名，所以为了避免冲突，这里用小写的router。
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// 为什么同包内可以直接调用的newRouter()方法，这里却需要用小写的newrouter()方法？
// 因为Go里的大写字母开头的函数名一般都是包外可见的，而小写字母开头的函数名一般都是包内可见的。
// 所以为了避免冲突，这里用小写的newrouter()方法。
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern

	_, ok := r.roots[method]

	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)

		for i, part := range parts {
			if part[0] == '*' && len(parts) > 1 {
				params[part[1:]] = strings.Join(searchParts[i:], "/")
				break
			}
			if part[0] == ':' {
				params[part[1:]] = searchParts[i]
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 page not found")
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}
