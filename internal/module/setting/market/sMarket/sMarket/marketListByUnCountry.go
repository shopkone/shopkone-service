package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/internal/module/setting/market/sMarket/sMarketCountry"
)

func (s *sMarket) MarketListByUnCountry() (out []mMarket.Market, err error) {
	// 获取市场列表
	if err = s.orm.Model(&mMarket.Market{}).Where("shop_id = ?", s.shopId).
		Select("id", "is_main", "name").Find(&out).Error; err != nil {
		return nil, err
	}

	// 查找市场下的国家
	marketIds := slice.Map(out, func(index int, item mMarket.Market) uint {
		return item.ID
	})
	countries, err := sMarketCountry.NewMarketCountry(s.orm, s.shopId).CountryList(marketIds)
	if err != nil {
		return nil, err
	}

	// 筛出没有绑定国家的市场
	out = slice.Filter(out, func(index int, i mMarket.Market) bool {
		_, has := slice.FindBy(countries, func(index int, item mMarket.MarketCountry) bool {
			return i.ID == item.MarketID
		})
		return !has
	})
	return out, err
}
