package sMarketProduct

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarketProduct) ProductList(marketID uint) ([]mMarket.MarketProduct, error) {
	var list []mMarket.MarketProduct
	return list, s.orm.Model(&mMarket.MarketProduct{}).
		Where("market_id = ?", marketID).
		Omit("created_at", "deleted_at", "updated_at", "shop_id").Find(&list).Error
}
