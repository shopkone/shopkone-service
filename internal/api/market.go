package api

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/base/resource/iResource"
	"shopkone-service/internal/module/base/resource/sResource"
	"shopkone-service/internal/module/setting/market/sMarket/sMarket"
	"shopkone-service/internal/module/setting/market/sMarket/sMarketProduct"
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
	c := ctx2.NewCtx(ctx)
	auth, err := c.GetAuth()
	shop := auth.Shop
	res, err = sMarket.NewMarket(sOrm.NewDb(&auth.Shop.ID), shop.ID).MarketList()
	if err != nil {
		return nil, err
	}
	countries := sResource.NewCountry().List()
	res = slice.Map(res, func(index int, item vo.MarketListRes) vo.MarketListRes {
		if item.IsMain {
			country, ok := slice.FindBy(countries, func(index int, country iResource.CountryListOut) bool {
				return country.Code == item.Name
			})
			if ok {
				item.Name = c.GetT(country.Name)
			}
		}
		return item
	})
	return res, err
}

func (a *aMarket) Info(ctx g.Ctx, req *vo.MarketInfoReq) (res vo.MarketInfoRes, err error) {
	c := ctx2.NewCtx(ctx)
	auth, err := c.GetAuth()
	shop := auth.Shop
	res, err = sMarket.NewMarket(sOrm.NewDb(&auth.Shop.ID), shop.ID).MarketInfo(req.ID)
	// 如果是主要市场，则翻译
	countries := sResource.NewCountry().List()
	if res.IsMain {
		country, ok := slice.FindBy(countries, func(index int, item iResource.CountryListOut) bool {
			return item.Code == res.Name
		})
		if ok {
			res.Name = c.GetT(country.Name)
		}
	}
	return res, err
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
	c := ctx2.NewCtx(ctx)
	auth, err := c.GetAuth()
	shop := auth.Shop
	return sMarket.NewMarket(sOrm.NewDb(&auth.Shop.ID), shop.ID).MarketOptions(c.GetT)
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

// 更新商品调整
func (a *aMarket) UpdateProduct(ctx g.Ctx, req *vo.MarketUpdateProductReq) (res vo.MarketUpdateProductRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		s := sMarketProduct.NewMarketProduct(tx, shop.ID)
		market, err := sMarket.NewMarket(tx, shop.ID).MarketSimple(req.MarketID)
		if err != nil {
			return err
		}
		productUpdateIn := sMarketProduct.ProductUpdateIn{
			Req:           req,
			StoreCurrency: shop.StoreCurrency,
			IsMain:        market.IsMain,
		}
		return s.ProductUpdate(productUpdateIn)
	})
	return res, err
}

// 获取商品调整
func (a *aMarket) GetProduct(ctx g.Ctx, req *vo.MarketGetProductReq) (res vo.MarketGetProductRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	shop := auth.Shop
	s := sMarketProduct.NewMarketProduct(sOrm.NewDb(&auth.Shop.ID), shop.ID)
	return s.GetPrice(req.MarketID, shop.StoreCurrency)
}

// 获取市场简
func (a *aMarket) Simple(ctx g.Ctx, req *vo.MarketSimpleReq) (res vo.MarketSimpleRes, err error) {
	c := ctx2.NewCtx(ctx)
	auth, err := c.GetAuth()
	shop := auth.Shop
	s := sMarket.NewMarket(sOrm.NewDb(&auth.Shop.ID), shop.ID)
	out, err := s.MarketSimple(req.ID)
	res.IsMain = out.IsMain
	res.Name = out.Name

	// 转换名称
	countries := sResource.NewCountry().List()
	if out.IsMain {
		country, ok := slice.FindBy(countries, func(index int, item iResource.CountryListOut) bool {
			return item.Code == out.Name
		})
		if ok {
			res.Name = c.GetT(country.Name)
		}
	}

	return res, err
}
