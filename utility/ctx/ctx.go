package ctx

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/text/gstr"
	"shopkone-service/internal/module/shop/user/mUser"
	"shopkone-service/utility/code"
	"shopkone-service/utility/middleware"
)

type ICtx interface {
	GetLanguage() string
	GetUa() string
	GetIp() string
	GetUser() mUser.User
}

type sCtx struct {
	ctx g.Ctx
}

func NewCtx(ctx g.Ctx) *sCtx {
	return &sCtx{gi18n.WithLanguage(ctx, "zh-CN")}
}

func (s *sCtx) GetLanguage() string {
	r := g.RequestFromCtx(s.ctx)
	if r == nil {
		return "en"
	}
	str := r.GetHeader("Accept-Language")
	if !gstr.Contains(str, ",") {
		return "en"
	}
	arr := gstr.Split(str, ",")
	return arr[0]
}

func (s *sCtx) GetUa() string {
	r := g.RequestFromCtx(s.ctx)
	if r == nil {
		return ""
	}
	return r.GetHeader("User-Agent")
}

func (s *sCtx) GetIp() string {
	return g.RequestFromCtx(s.ctx).GetClientIp()
}

func (s *sCtx) GetUser() (user mUser.User, err error) {
	r := g.RequestFromCtx(s.ctx)
	if r == nil {
		return user, code.AuthError
	}
	if err = r.GetCtxVar("shopkone_user").Scan(&user); err != nil {
		return user, code.AuthError
	}
	if user.ID <= 0 {
		return user, code.AuthError
	}
	return user, nil
}

func (s *sCtx) GetAuth() (auth middleware.Auth, err error) {
	r := g.RequestFromCtx(s.ctx)
	if r == nil {
		return auth, code.AuthError
	}
	if err = r.GetCtxVar("shopkone_auth").Scan(&auth); err != nil {
		return auth, code.AuthError
	}
	if auth.User.ID <= 0 {
		return auth, code.AuthError
	}
	if auth.Shop.ID <= 0 {
		return auth, code.AuthError
	}
	if auth.Staff.ID <= 0 {
		return auth, code.AuthError
	}
	return auth, nil
}

func (s *sCtx) GetT(key string) string {
	return gi18n.T(s.ctx, key)
}
