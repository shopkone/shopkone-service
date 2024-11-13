package sMarketCountry

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarketCountry) CountryCreate(codes []string, marketID uint) (err error) {
	if len(codes) == 0 {
		return err
	}

	// 删除原有的
	if err = s.CountryRemove(codes); err != nil {
		return err
	}
	// 创建新的
	data := slice.Map(codes, func(_ int, code string) mMarket.MarketCountry {
		i := mMarket.MarketCountry{
			CountryCode: code,
			MarketID:    marketID,
		}
		i.ShopId = s.shopId
		return i
	})
	return s.orm.Create(&data).Error
}
