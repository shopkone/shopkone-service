package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/code"
)

func (s *sMarket) LanguageCheck() (err error) {
	markets, err := s.MarketOptions()
	if err != nil {
		return err
	}
	marketIds := slice.Map(markets, func(index int, item vo.MarketOptionsRes) uint {
		return item.Value
	})

	// 获取语言绑定关系
	languages, err := s.LanguagesByMarketIds(marketIds)
	if err != nil {
		return err
	}

	// 市场必须要有一个语言
	isAllHas := slice.Every(marketIds, func(index int, item uint) bool {
		_, exist := slice.FindBy(languages, func(index int, i mMarket.MarketLanguage) bool {
			return i.MarketID == item
		})
		return exist
	})
	if !isAllHas {
		return code.MarketMustLanguage
	}

	return err
}
