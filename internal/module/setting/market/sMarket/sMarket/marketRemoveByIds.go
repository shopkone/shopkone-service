package sMarket

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarket) MarketRemoveByIds(ids []uint) (err error) {
	// 删除这些市场
	return s.orm.Where("shop_id = ? AND id IN ?", s.shopId, ids).
		Delete(&mMarket.Market{}).Error
}
