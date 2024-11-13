package sMarketCountry

import "shopkone-service/internal/module/setting/market/mMarket"

func (s *sMarketCountry) CountryRemove(codes []string) (err error) {
	return s.orm.
		Where("country_code in (?) AND shop_id = ?", codes, s.shopId).
		Unscoped().Delete(&mMarket.MarketCountry{}).Error
}
