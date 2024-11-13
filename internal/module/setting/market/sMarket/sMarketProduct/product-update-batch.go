package sMarketProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/handle"
)

func (s *sMarketProduct) ProductUpdateBatch(products, oldProducts []mMarket.MarketProduct) (err error) {
	products = slice.Filter(products, func(index int, newItem mMarket.MarketProduct) bool {
		oldItem, ok := slice.FindBy(oldProducts, func(index int, oldItem mMarket.MarketProduct) bool {
			return oldItem.ID == newItem.ID
		})
		if !ok {
			return false
		}
		return s.ProductIsChange(oldItem, newItem)
	})
	if len(products) == 0 {
		return err
	}
	products = slice.Map(products, func(index int, item mMarket.MarketProduct) mMarket.MarketProduct {
		item.CanCreateId = true
		return item
	})
	batchIn := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"exclude", "fixed", "type"},
	}
	return handle.BatchUpdateById(batchIn, &products)
}
