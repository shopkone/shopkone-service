package sMarketProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarketProduct) ProductCreate(products []mMarket.MarketProduct) (err error) {
	if len(products) == 0 {
		return err
	}
	products = slice.Map(products, func(index int, item mMarket.MarketProduct) mMarket.MarketProduct {
		item.ID = 0
		return item
	})
	return s.orm.Create(products).Error
}
