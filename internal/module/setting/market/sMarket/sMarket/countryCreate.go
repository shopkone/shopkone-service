package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarket) CountryCreate(codes []string, marketID uint) (err error) {
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
