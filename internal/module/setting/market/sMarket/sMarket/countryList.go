package sMarket

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarket) CountryList(marketIds []uint) (out []mMarket.MarketCountry, err error) {
	if err = s.orm.Model(&mMarket.MarketCountry{}).
		Where("market_id IN (?) AND shop_id = ?", marketIds, s.shopId).
		Omit("shop_id", "created_at", "updated_at", "deleted_at").Find(&out).Error; err != nil {
		return nil, err
	}

	return out, err
}
