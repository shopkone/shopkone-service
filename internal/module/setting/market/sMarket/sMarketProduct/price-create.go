package sMarketProduct

import "shopkone-service/internal/module/setting/market/mMarket"

type MarketPriceCreateIn struct {
	AdjustPercent float64
	AdjustType    mMarket.PriceAdjustmentType
	CurrencyCode  string
	MarketID      uint
}

func (s *sMarketProduct) PriceCreate(in MarketPriceCreateIn) (err error) {
	var price mMarket.MarketPrice
	price.MarketID = in.MarketID
	price.ShopId = s.shopId
	price.CurrencyCode = in.CurrencyCode
	price.AdjustPercent = in.AdjustPercent
	price.AdjustType = in.AdjustType
	return s.orm.Create(&price).Error
}
