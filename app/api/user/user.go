package user

import "github.com/gogf/gf/net/ghttp"

func Login(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}