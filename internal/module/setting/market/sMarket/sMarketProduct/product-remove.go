package sMarketProduct

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarketProduct) ProductRemove(ids []uint) (err error) {
	if len(ids) == 0 {
		return nil
	}
	return s.orm.Model(&mMarket.MarketProduct{}).
		Where("shop_id = ?", s.shopId).
		Where("id IN ?", ids).
		Unscoped().
		Delete(&mMarket.MarketProduct{}).Error
}
