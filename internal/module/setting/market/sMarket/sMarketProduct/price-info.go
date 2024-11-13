package sMarketProduct

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarketProduct) PriceInfo(marketId uint) (price mMarket.MarketPrice, err error) {
	return price, s.orm.Model(&price).
		Where("market_id = ?", marketId).
		Select("id", "adjust_percent", "adjust_type", "currency_code").
		First(&price).Error
}
