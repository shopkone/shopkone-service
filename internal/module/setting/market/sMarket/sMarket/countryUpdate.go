package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarket) CountryUpdate(codes []string, marketId uint) (err error) {
	// 获取该market下面的country
	oldCountries, err := s.CountryList([]uint{marketId})
	if err != nil {
		return err
	}
	oldCodes := slice.Map(oldCountries, func(_ int, item mMarket.MarketCountry) string {
		return item.CountryCode
	})

	// 找出要删除的codes
	deleteCodes := slice.Filter(oldCodes, func(_ int, item string) bool {
		return slice.Contain(codes, item) == false
	})
	if len(deleteCodes) > 0 {
		if err = s.orm.Model(&mMarket.MarketCountry{}).
			Where("market_id = ? AND country_code IN ? AND shop_id = ?", marketId, deleteCodes, s.shopId).
			Unscoped().Delete(&mMarket.MarketCountry{}).Error; err != nil {
			return err
		}
		// 解绑市场语言
	}
	// 找出要新增的codes
	addCodes := slice.Filter(codes, func(_ int, item string) bool {
		return slice.Contain(oldCodes, item) == false
	})
	return s.CountryCreate(addCodes, marketId)
}
