package gee

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

type HandlerFunc func(c *Context)

type RouterGroup struct {
	prefix    string
	middlware []HandlerFunc
	gee       *Gee //所有组都指向唯一的个一Gee
	parent    *RouterGroup
}

func (g *RouterGroup) Get(uri string, handler HandlerFunc) {
	g.gee.router.addRouter("GET", g.prefix+uri, handler)
}

func (g *RouterGroup) Post(uri string, handler HandlerFunc) {
	g.gee.router.addRouter("POST", g.prefix+uri, handler)
}

func (g *RouterGroup) Use(handlers ...HandlerFunc) {
	g.middlware = append(g.middlware, handlers...)
}

func (g *RouterGroup) createStaticHandler(dirPath string) HandlerFunc {
	fs := http.Dir(dirPath)
	absolutePath := path.Join(g.prefix, dirPath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param(WildUri)
		if _, err := fs.Open(file); err != nil {
			fmt.Println("file server error:", err)
			handle404(c)
			return
		}
		//fileServer.ServeHTTP will using c.Req.URL.Path to read data
		c.Req.URL.Path = absolutePath + file
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (g *RouterGroup) Static(uri, dirPath string) {
	handler := g.createStaticHandler(dirPath)
	realUri := uri + WildUri
	g.Get(realUri, handler)
}

type Gee struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup
}

func NewGee() *Gee {
	g := &Gee{router: newRouter()}
	group := &RouterGroup{gee: g}
	g.RouterGroup = group
	g.groups = append(g.groups, group)
	return g
}

func Default() *Gee {
	g := NewGee()
	g.Use(Recovery())
	return g
}

func (g *Gee) Group(prefix string) *RouterGroup {
	return g.RouterGroup.AddChildGroup(prefix)
}

func (g *RouterGroup) AddChildGroup(prefix string) *RouterGroup {
	group := &RouterGroup{prefix: g.prefix + prefix, gee: g.gee}
	g.gee.groups = append(g.gee.groups, group)
	group.parent = g
	return group
}

func (g *Gee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	//由于数据流都是通过Gee引擎来进行处理，这里ServeHTTP也只能是Gee的方法，
	//同时中间件以组为单位进行创建(Gee也包含超级组)，因此这里要进行handler的一个遍历
	//赋值方法到实际的Context中
	for _, group := range g.groups {
		if strings.HasPrefix(req.RequestURI, group.prefix) {
			c.handlers = append(c.handlers, group.middlware...)
		}
	}
	g.router.handle(c)
}

func (g *Gee) Run() error {
	return http.ListenAndServe("127.0.0.1:80", g)
}
