package middleware

import (
	"github.com/gogf/gf/v2/encoding/gcompress"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 压缩
func Compress(r *ghttp.Request) {
	r.Middleware.Next()
	ghttp.MiddlewareHandlerResponse(r)
	ret, err := gcompress.Gzip(r.Response.Buffer())
	if err != nil {
		r.Response.WriteStatusExit(500)
		return
	}
	r.Response.ClearBuffer()
	r.Response.Header().Set("Content-Encoding", "gzip")
	r.Response.Write(ret)
}
