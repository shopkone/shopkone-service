package sMarket

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarket) UnbindByMarketIds(ids []uint) (err error) {
	return s.orm.Model(&mMarket.MarketLanguage{}).
		Where("shop_id = ? AND market_id IN ?", s.shopId, ids).
		Delete(&mMarket.MarketLanguage{}).Error
}
