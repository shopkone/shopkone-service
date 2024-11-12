package sMarket

import (
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarket) ProductsList(marketID uint) (res []mMarket.MarketProduct, err error) {
	if err = s.orm.Model(&res).
		Where("market_id = ?", marketID).
		Omit("created_at", "shop_id", "deleted_at", "updated_at").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, err
}
