package sMarket

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarket) CountryRemove(codes []string) (err error) {
	return s.orm.
		Where("country_code in (?) AND shop_id = ?", codes, s.shopId).
		Delete(&mMarket.MarketCountry{}).Error
}
