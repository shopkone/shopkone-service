package sMarket

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarket) GetMarketsByLanguageIds(languageIds []uint) (markets []mMarket.MarketLanguage, err error) {
	if err = s.orm.Model(&markets).Where("language_id in (?) AND shop_id = ?", languageIds, s.shopId).
		Select("language_id", "market_id", "id", "is_default").Find(&markets).Error; err != nil {
		return
	}
	return markets, err
}
