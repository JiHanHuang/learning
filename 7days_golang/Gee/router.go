package gee

import "fmt"

type Router struct {
	roots  *node
	router map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{roots: NewNode(), router: make(map[string]HandlerFunc)}
}

func (r *Router) addRouter(method, uri string, hander HandlerFunc) {
	r.roots.Add(uri)
	key := method + "-" + uri
	r.router[key] = hander
}

func (r *Router) handle(c *Context) {

	node, k := r.roots.Get(c.Req.RequestURI)
	if k == "" || node == nil {
		errorHandle(404, c)
	} else {
		c.n = node
		key := c.Req.Method + "-" + k
		if handler, ok := r.router[key]; ok {
			c.handlers = append(c.handlers, handler)
		} else {
			errorHandle(404, c)
		}
	}
	c.Next()
}

func errorHandle(code int, c *Context) {
	c.handlers = make([]HandlerFunc, 0, 1)
	switch code {
	case 404:
		c.handlers = append(c.handlers, handle404)
	default:
		c.handlers = append(c.handlers, func(c *Context) {
			fmt.Fprintf(c.Writer, "500 INTERNAL ERROR\n")
		})
	}
}

func handle404(c *Context) {
	fmt.Fprintf(c.Writer, "404 NOT FOUND: %s\n", c.Req.URL.Path)
	fmt.Fprintf(c.Writer, "Req: \n Url:%s\n Header:%v\n", c.Req.URL.Path, c.Req.Header)
}
