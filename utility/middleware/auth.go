package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"shopkone-service/internal/module/shop/shop/mShop"
	"shopkone-service/internal/module/shop/shop/sShop"
	"shopkone-service/internal/module/shop/staff/mStaff"
	"shopkone-service/internal/module/shop/staff/sStaff"
	"shopkone-service/internal/module/shop/user/mUser"
	"shopkone-service/internal/module/shop/user/sUser"
	"strings"
)

type Auth struct {
	Shop  mShop.Shop
	User  mUser.User
	Staff mStaff.Staff
}

func AuthMiddleware(r *ghttp.Request) {
	// 校验token
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
	// 获取user
	user, err := sUser.NewUserCache().GetUserLoginCache(arr[1])
	if err != nil || user.ID == 0 {
		UnAuthorized(r)
	}
	// 校验shopUuid
	uuid := r.Header.Get("x-shopid")
	if len(uuid) != 64 {
		UnAuthorized(r)
	}
	// 获取shop
	shop, err := sShop.NewShopCache().GetShopCache(uuid)
	if err != nil || shop.ID == 0 {
		UnAuthorized(r)
	}
	// 获取staff
	staff, err := sStaff.NewStaffCache().GetCache(shop.ID, user.ID)
	if err != nil || staff.ID == 0 {
		UnAuthorized(r)
	}
	auth := Auth{
		Shop:  shop,
		User:  user,
		Staff: staff,
	}
	r.SetCtxVar("shopkone_auth", auth)
	r.Middleware.Next()
}
