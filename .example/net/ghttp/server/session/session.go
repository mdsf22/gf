package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		g.GET("/set", func(r *ghttp.Request) {
			r.Session.Set("time", gtime.Second())
			r.Response.Write("ok")
		})
		g.GET("/get", func(r *ghttp.Request) {
			r.Response.WriteJson(r.Session.Map())
		})
		g.GET("/clear", func(r *ghttp.Request) {
			r.Session.Clear()
		})
	})
	s.SetPort(8199)
	s.Run()
}
