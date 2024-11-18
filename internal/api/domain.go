package api

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/base/resource/iResource"
	"shopkone-service/internal/module/base/resource/sResource"
	"shopkone-service/internal/module/setting/domains/sDomain/sBlockCountry"
	"shopkone-service/internal/module/setting/domains/sDomain/sDomain"
	ctx2 "shopkone-service/utility/ctx"

	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
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
	return sDomain.NewDomain(sOrm.NewDb(&auth.Shop.ID), shop.ID).DomainList(req)
}

func (a *aDomain) PreCheck(ctx g.Ctx, req *vo.DomainPreCheckReq) (res []vo.DomainPreCheckRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
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
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		in := sDomain.ConnectCheckIn{
			Domain:     req.Domain,
			IsShopkone: false,
		}
		res.IsConnect = true
		return sDomain.NewDomain(tx, shop.ID).ConnectCheck(in)
	})
	return res, err
}

// 更新屏蔽国家
func (a *aDomain) BlockCountryUpdate(ctx g.Ctx, req *vo.DomainBlockCountryUpdateReq) (res vo.DomainBlockCountryUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	err = sOrm.NewDb(&shop.ID).Transaction(func(tx *gorm.DB) error {
		s := sBlockCountry.NewBlockCountry(tx, shop.ID)
		return s.Update(req.Codes)
	})
	return res, err
}

// 获取屏蔽国家
func (a *aDomain) BlockCountryList(ctx g.Ctx, req *vo.DomainBlockCountryListReq) (res []vo.DomainBlockCountryListRes, err error) {
	c := ctx2.NewCtx(ctx)
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}
	shop := auth.Shop
	countryCodes, err := sBlockCountry.NewBlockCountry(sOrm.NewDb(&shop.ID), shop.ID).List()
	if err != nil {
		return nil, err
	}
	countries := sResource.NewCountry().List()
	t := c.GetT
	res = slice.Map(countryCodes, func(index int, code string) vo.DomainBlockCountryListRes {
		cy, ok := slice.FindBy(countries, func(index int, c iResource.CountryListOut) bool {
			return c.Code == code
		})
		if !ok {
			return vo.DomainBlockCountryListRes{}
		}
		i := vo.DomainBlockCountryListRes{
			Code: code,
			Name: t(cy.Name),
		}
		return i
	})
	return res, err
}
