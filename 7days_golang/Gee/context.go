package gee

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	n      *node
	index  int
	//每一个请求都会有一些列动作，比如数据处理，数据格式校验，日志，错误处理等，都称为handlers，
	//每个请求不同，遇到不同情况，动态修改Context里的handler
	//当然，Group里的中间件也会根据需求增加到这里面
	handlers []HandlerFunc
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		index:  -1,
	}
}

func (c *Context) SetHeader(k, v string) {
	c.Writer.Header().Set(k, v)
}
func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)

}

func (c *Context) Json(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
	//不能把自增放到这里，否则handlers调用了Next, for循环里就会出现无限递归
	//c.index++
}

func (c *Context) Param(key string) string {
	switch key {
	case WildUri:
		return c.n.value
	}
	return "'"
}
