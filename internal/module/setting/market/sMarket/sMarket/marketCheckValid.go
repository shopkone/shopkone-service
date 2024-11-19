package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/code"
)

func (s *sMarket) MarketCheckValid() (err error) {
	markets, err := s.MarketOptions(nil)
	if err != nil {
		return err
	}
	marketIds := slice.Map(markets, func(index int, market vo.MarketOptionsRes) uint {
		return market.Value
	})

	var count int64
	if err = s.orm.Model(&mMarket.Market{}).
		Where("shop_id = ? AND id IN ?", s.shopId, marketIds).
		Count(&count).Error; err != nil {
		return err
	}
	if count != int64(len(marketIds)) {
		return code.MarketValid
	}
	return err
}
