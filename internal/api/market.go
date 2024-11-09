package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/setting/language/sLanguage"
	"shopkone-service/internal/module/setting/market/sMarket/sMarket"
	"shopkone-service/internal/module/setting/market/sMarket/sMarketLanguage"
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

func (a *aMarket) Options(ctx g.Ctx, req *vo.MarketOptionsReq) (res []vo.MarketOptionsRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	return sMarket.NewMarket(sOrm.NewDb(), shop.ID).MarketOptions()
}

// 更新市场域名设置
func (a *aMarket) UpDomain(ctx g.Ctx, req *vo.MarketUpDomainReq) (res vo.MarketUpDomainRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		s := sMarket.NewMarket(tx, shop.ID)
		return s.MarketUpdateDomain(*req)
	})
	return res, err
}

func (a *aMarket) BindByLangId(ctx g.Ctx, req *vo.BindLangByLangIdReq) (res vo.BindLangByLangIdRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		// 过滤一下，只绑定是这个店铺的市场
		req.MarketIDs, err = sMarket.NewMarket(tx, shop.ID).MarketFilterIds(req.MarketIDs)
		if err != nil {
			return err
		}
		// 根据语言绑定市场
		s := sMarketLanguage.NewMarketLanguage(tx, shop.ID)
		return s.BindByLanguageId(req)
	})
	return res, err
}

func (a *aMarket) BindByMarketId(ctx g.Ctx, req *vo.BindLangByMarketIdReq) (res vo.BindLangByMarketIdRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		// 过滤一下，只绑定是这个店铺的语言
		req.LanguageIDs, err = sLanguage.NewLanguage(tx, shop.ID).LanguageFilter(req.LanguageIDs)
		if err != nil {
			return err
		}
		// 根据市场绑定语言
		s := sMarketLanguage.NewMarketLanguage(tx, shop.ID)
		return s.BindLangByMarketId(req)
	})
	return res, err
}
