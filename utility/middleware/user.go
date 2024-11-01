package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"shopkone-service/internal/module/shop/user/sUser"
	"strings"
)

func UserMiddleware(r *ghttp.Request) {
	token := r.Header.Get("authorization")
	if token == "" {
		UnAuthorized(r)
	}
	arr := strings.Split(token, " ")
	if len(arr) != 2 {
		UnAuthorized(r)
	}
	if arr[0] != "Bearer" {
		UnAuthorized(r)
	}
	if len(arr[1]) != 64 {
		UnAuthorized(r)
	}
	user, err := sUser.NewUserCache().GetUserLoginCache(arr[1])
	if err != nil || user.ID == 0 {
		UnAuthorized(r)
	}
	r.SetCtxVar("shopkone_user", user)
	r.Middleware.Next()
}

func UnAuthorized(r *ghttp.Request) {
	r.Response.WriteStatusExit(401, "Unauthorized")
	r.ExitAll()
}
