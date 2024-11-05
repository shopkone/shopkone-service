package sMarketLanguage

import (
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarketLanguage) LanguageBind(languageId uint, marketId uint) (err error) {
	var data mMarket.MarketLanguage
	data.ShopId = s.shopId
	data.LanguageID = languageId
	data.MarketID = marketId
	return s.orm.Create(&data).Error
}
