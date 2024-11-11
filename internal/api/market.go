package api

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/setting/market/sMarket/sMarket"
	ctx2 "shopkone-service/utility/ctx"

	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
)

type aMarket struct {
}

func NewMarketApi() *aMarket {
	return &aMarket{}
}

func (a *aMarket) Create(ctx g.Ctx, req *vo.MarketCreateReq) (res vo.MarketCreateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		res, err = sMarket.NewMarket(tx, shop.ID).MarketCreate(*req)
		return err
	})
	return res, err
}

func (a *aMarket) List(ctx g.Ctx, req *vo.MarketListReq) (res []vo.MarketListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sMarket.NewMarket(sOrm.NewDb(&auth.Shop.ID), shop.ID).MarketList()
}

func (a *aMarket) Info(ctx g.Ctx, req *vo.MarketInfoReq) (res vo.MarketInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sMarket.NewMarket(sOrm.NewDb(&auth.Shop.ID), shop.ID).MarketInfo(req.ID)
}

func (a *aMarket) Update(ctx g.Ctx, req *vo.MarketUpdateReq) (res vo.MarketUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		res, err = sMarket.NewMarket(tx, shop.ID).MarketUpdate(*req)
		return err
	})
	return res, err
}

func (a *aMarket) Options(ctx g.Ctx, req *vo.MarketOptionsReq) (res []vo.MarketOptionsRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sMarket.NewMarket(sOrm.NewDb(&auth.Shop.ID), shop.ID).MarketOptions()
}

// 更新市场域名设置
func (a *aMarket) UpDomain(ctx g.Ctx, req *vo.MarketUpDomainReq) (res vo.MarketUpDomainRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		s := sMarket.NewMarket(tx, shop.ID)
		return s.MarketUpdateDomain(*req)
	})
	return res, err
}

// 更新市场语言
func (a *aMarket) UpdateLang(ctx g.Ctx, req *vo.MarketUpdateLangReq) (res vo.MarketUpdateLangRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		s := sMarket.NewMarket(tx, shop.ID)
		return s.MarketUpdateLang(*req)
	})
	return res, err
}

// 根据语言id更新市场
func (a *aMarket) UpdateLangByLangID(ctx g.Ctx, req *vo.MarketUpdateLangByLangIDReq) (res vo.MarketUpdateLangByLangIDRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		s := sMarket.NewMarket(tx, shop.ID)
		return s.MarketUpdateLangByLangID(req)
	})
	return res, err
}
