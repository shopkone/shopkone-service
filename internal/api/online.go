package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/online/mav/sNav"
	ctx2 "shopkone-service/utility/ctx"
)

type aOnlineApi struct {
}

func NewOnlineApi() *aOnlineApi {
	return &aOnlineApi{}
}

func (a *aOnlineApi) NavList(ctx g.Ctx, req *vo.OnlineNavListReq) (res []vo.OnlineNavListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sNav.NewNav(sOrm.NewDb(&auth.Shop.ID), shop.ID).NavList()
}
