package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/setting/market/sMarket/sMarket"
	ctx2 "shopkone-service/utility/ctx"
)

type aMarket struct {
}

func NewMarketApi() *aMarket {
	return &aMarket{}
}

func (a *aMarket) Create(ctx g.Ctx, req *vo.MarketCreateReq) (res vo.MarketCreateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		res, err = sMarket.NewMarket(tx, shop.ID).MarketCreate(*req)
		return err
	})
	return res, err
}

func (a *aMarket) List(ctx g.Ctx, req *vo.MarketListReq) (res []vo.MarketListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sMarket.NewMarket(sOrm.NewDb(), shop.ID).MarketList()
}

func (a *aMarket) Info(ctx g.Ctx, req *vo.MarketInfoReq) (res vo.MarketInfoRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sMarket.NewMarket(sOrm.NewDb(), shop.ID).MarketInfo(req.ID)
}

func (a *aMarket) Update(ctx g.Ctx, req *vo.MarketUpdateReq) (res vo.MarketUpdateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		res, err = sMarket.NewMarket(tx, shop.ID).MarketUpdate(*req)
		return err
	})
	return res, err
}
