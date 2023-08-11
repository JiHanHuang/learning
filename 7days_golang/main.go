package main

import (
	"fmt"
	"log"

	"gee.com"
)

func main() {
	geeMain()
}

func geeMain() {
	r := gee.Default()
	r.Get("/url", func(c *gee.Context) {
		fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Req.URL.Path)
	})
	r.Get("/hello", func(c *gee.Context) {
		for k, v := range c.Req.Header {
			fmt.Fprintf(c.Writer, "Header[%q] = %q\n", k, v)
		}
	})
	obj := map[string]string{
		"name": "li",
		"age":  "11",
	}
	r.Get("/login", func(c *gee.Context) {
		c.Json(200, obj)
	})
	r.Get("/hi/:name", func(c *gee.Context) {
		fmt.Fprintf(c.Writer, "hi = %q\n", c.Req.URL.Path)
	})

	r.Static("/file", "/root/code/7days_golang/")

	r.Get("/panic", func(c *gee.Context) {
		var a []string
		a[0] = "hello"
	})

	//group
	v1 := r.Group("/v1")
	v1.Get("/hello", func(c *gee.Context) {
		fmt.Fprintf(c.Writer, "group v1 hello = %q\n", c.Req.URL.Path)
	})

	//middler
	v2 := r.Group("/v2")
	v2.Use(func(c *gee.Context) {
		fmt.Fprintf(c.Writer, "before v2\n")
		c.Next()
		fmt.Fprintf(c.Writer, "end v2\n")
	})
	v2.Get("/hello", func(c *gee.Context) {
		fmt.Fprintf(c.Writer, "group v2 hello = %q\n", c.Req.URL.Path)
	})
	log.Println("Gee listen...")
	log.Fatal(r.Run())
}
