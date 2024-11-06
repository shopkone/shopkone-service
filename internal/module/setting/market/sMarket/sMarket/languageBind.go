package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/market/mMarket"
)

type LanguageBindIn struct {
	LanguageId uint   `json:"language_id"`
	MarketIds  []uint `json:"market_ids"`
}

// 语言的维度绑定
func (s *sMarket) LanguageBind(in LanguageBindIn) (err error) {
	if in.LanguageId == 0 || len(in.MarketIds) == 0 {
		return err
	}

	// 绑定
	data := slice.Map(in.MarketIds, func(index int, item uint) mMarket.MarketLanguage {
		i := mMarket.MarketLanguage{
			LanguageID: in.LanguageId,
			MarketID:   item,
		}
		i.ShopId = s.shopId
		return i
	})
	if err = s.orm.Create(&data).Error; err != nil {
		return err
	}

	// 后置校验
	return s.LanguageCheck()
}
