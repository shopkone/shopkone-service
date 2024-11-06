package sMarket

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarket) LanguagesByMarketIds(marketIds []uint) (out []mMarket.MarketLanguage, err error) {
	if err = s.orm.Model(&out).Where("market_id in (?)", marketIds).
		Omit("shop_id", "created_at", "updated_at", "deleted_at").Find(&out).Error; err != nil {
		return out, err
	}
	return out, err
}
