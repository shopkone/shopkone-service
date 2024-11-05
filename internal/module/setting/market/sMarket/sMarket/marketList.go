package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarket) MarketList() (out []vo.MarketListRes, err error) {
	var list []mMarket.Market
	// 获取市场列表
	if err = s.orm.Model(&mMarket.Market{}).Where("shop_id = ?", s.shopId).
		Omit("shop_id", "created_at", "deleted_at", "updated_at").Find(&list).Error; err != nil {
		return nil, err
	}

	// 获取市场下的国家
	marketIds := slice.Map(list, func(index int, item mMarket.Market) uint {
		return item.ID
	})
	codes, err := s.CountryList(marketIds)
	if err != nil {
		return nil, err
	}
	out = slice.Map(list, func(index int, market mMarket.Market) vo.MarketListRes {
		currentCodes := slice.Filter(codes, func(index int, code mMarket.MarketCountry) bool {
			return code.MarketID == market.ID
		})
		i := vo.MarketListRes{}
		i.Name = market.Name
		i.IsMain = market.IsMain
		i.ID = market.ID
		i.CountryCodes = slice.Map(currentCodes, func(index int, item mMarket.MarketCountry) string {
			return item.CountryCode
		})
		return i
	})
	return out, err
}
