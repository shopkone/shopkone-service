package api

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/orm/sOrm"
	"shopkone-service/internal/module/setting/language/mLanguage"
	"shopkone-service/internal/module/setting/language/sLanguage"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/internal/module/setting/market/sMarket/sMarket"
	ctx2 "shopkone-service/utility/ctx"
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
	orm := sOrm.NewDb()

	// 获取语言
	languages, err := sLanguage.NewLanguage(orm, shop.ID).LanguageList()
	if err != nil {
		return nil, err
	}

	// 根据语言获取marketIds
	languageIds := slice.Map(languages, func(index int, item mLanguage.Language) uint {
		return item.ID
	})
	marketLanguages, err := sMarket.NewMarket(orm, shop.ID).BindListByLangIds(languageIds)
	if err != nil {
		return nil, err
	}

	// 组装数据
	return slice.Map(languages, func(index int, item mLanguage.Language) vo.LanguageListRes {
		currentMarketLanguages := slice.Filter(marketLanguages, func(index int, m mMarket.MarketLanguage) bool {
			return m.LanguageID == item.ID
		})
		markets := slice.Map(currentMarketLanguages, func(index int, item mMarket.MarketLanguage) vo.LanguageMarket {
			i := vo.LanguageMarket{
				MarketID:  item.MarketID,
				IsDefault: item.IsDefault,
			}
			return i
		})
		return vo.LanguageListRes{
			ID:        item.ID,
			IsDefault: item.IsDefault,
			Language:  item.Code,
			Markets:   markets,
		}
	}), nil
}

func (a *aLanguage) Create(ctx g.Ctx, req *vo.LanguageCreateReq) (res vo.LanguageCreateRes, err error) {
	auth, err := ctx2.NewCtx(ctx).GetAuth()
	if err != nil {
		return res, err
	}

	shop := auth.Shop
	orm := sOrm.NewDb()

	err = sOrm.NewDb().Transaction(func(tx *gorm.DB) error {
		_, err = sLanguage.NewLanguage(orm, shop.ID).LanguageCreate(req.Codes, false)
		return err
	})

	return res, err
}
