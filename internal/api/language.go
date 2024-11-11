package api

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/setting/language/mLanguage"
	"shopkone-service/internal/module/setting/language/sLanguage"
	"shopkone-service/internal/module/setting/market/sMarket/sMarket"
	ctx2 "shopkone-service/utility/ctx"

	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
)

type aLanguage struct {
}

func NewLanguageApi() *aLanguage {
	return &aLanguage{}
}

func (a *aLanguage) List(ctx g.Ctx, req *vo.LanguageListReq) (res []vo.LanguageListRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}

	shop := auth.Shop
	orm := sOrm.NewDb(&auth.Shop.ID)

	// 获取语言
	languages, err := sLanguage.NewLanguage(orm, shop.ID).LanguageList()
	if err != nil {
		return nil, err
	}

	// 获取市场列表
	marketsOptions, err := sMarket.NewMarket(orm, shop.ID).MarketOptions()
	if err != nil {
		return nil, err
	}

	// 组装数据
	return slice.Map(languages, func(index int, lang mLanguage.Language) vo.LanguageListRes {
		currentMarket := slice.Filter(marketsOptions, func(index int, market vo.MarketOptionsRes) bool {
			_, ok := slice.FindBy(market.LanguageIds, func(index int, langId uint) bool {
				return langId == lang.ID
			})
			return ok
		})
		MarketNames := slice.Map(currentMarket, func(index int, item vo.MarketOptionsRes) string {
			return item.Label
		})
		i := vo.LanguageListRes{}
		i.Language = lang.Code
		i.ID = lang.ID
		i.IsDefault = lang.IsDefault
		i.MarketNames = MarketNames
		return i
	}), nil
}

func (a *aLanguage) Create(ctx g.Ctx, req *vo.LanguageCreateReq) (res vo.LanguageCreateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}

	shop := auth.Shop
	orm := sOrm.NewDb(&auth.Shop.ID)

	err = sOrm.NewDb(&auth.Shop.ID).Transaction(func(tx *gorm.DB) error {
		_, err = sLanguage.NewLanguage(orm, shop.ID).LanguageCreate(req.Codes, false)
		return err
	})

	return res, err
}
