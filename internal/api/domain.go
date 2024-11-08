package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/setting/domains/sDomain/sDomain"
	ctx2 "shopkone-service/utility/ctx"
)

type aDomain struct {
}

func NewDomainApi() *aDomain {
	return &aDomain{}
}

func (a *aDomain) List(ctx g.Ctx, req *vo.DomainListReq) (res []vo.DomainListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return nil, err
	}
	shop := auth.Shop
	return sDomain.NewDomain(sOrm.NewDb(), shop.ID).DomainList(req)
}

func (a *aDomain) PreCheck(ctx g.Ctx, req *vo.DomainPreCheckReq) (res []vo.DomainPreCheckRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		out, err := sDomain.NewDomain(tx, shop.ID).PreCheck(req.Domain)
		res = []vo.DomainPreCheckRes{
			{
				Type:  "A",
				Host:  "@",
				Value: out.BindIp,
			},
			{
				Type:  "CNAME",
				Host:  "www",
				Value: out.BindDomain,
			},
		}
		return err
	})
	return res, err
}

func (a *aDomain) ConnectCheck(ctx g.Ctx, req *vo.DomainConnectCheckReq) (res vo.DomainConnectCheckRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		in := sDomain.ConnectCheckIn{
			Domain:     req.Domain,
			IsShopkone: false,
		}
		res.IsConnect = true
		return sDomain.NewDomain(tx, shop.ID).ConnectCheck(in)
	})
	return res, err
}
