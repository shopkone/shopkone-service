package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/online/mav/mNav"
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

func (a *aOnlineApi) NavInfo(ctx g.Ctx, req *vo.OnlineNavInfoReq) (res vo.OnlineNavInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sNav.NewNav(sOrm.NewDb(&auth.Shop.ID), shop.ID).NavInfo(*req)
}

func (a *aOnlineApi) NavUpdate(ctx g.Ctx, req *vo.OnlineNavUpdateReq) (res vo.OnlineNavUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sNav.NewNav(sOrm.NewDb(&auth.Shop.ID), shop.ID).NavUpdate(*req)
}

func (a *aOnlineApi) NavCreate(ctx g.Ctx, req *vo.OnlineNavCreateReq) (res vo.OnlineNavCreateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	in := mNav.Nav{
		Handle: req.Handle,
		Links:  sNav.FillLink(req.Links),
		Title:  req.Title,
	}
	res.ID, err = sNav.NewNav(sOrm.NewDb(&auth.Shop.ID), shop.ID).NavCreate(in)
	return res, err
}
