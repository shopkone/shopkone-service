package sMarketProduct

import "shopkone-service/internal/module/setting/market/mMarket"

type MarketPriceUpdateIn struct {
	MarketID      uint
	CurrencyCode  string
	AdjustPercent float64
	AdjustType    mMarket.PriceAdjustmentType
}

func (s *sMarketProduct) PriceUpdate(in MarketPriceUpdateIn) (err error) {
	updateIn := mMarket.MarketPrice{}
	updateIn.AdjustPercent = in.AdjustPercent
	updateIn.AdjustType = in.AdjustType
	updateIn.CurrencyCode = in.CurrencyCode
	query := s.orm.Model(&updateIn).Where("market_id = ?", in.MarketID)
	query = query.Select("adjust_percent", "adjust_type", "currency_code")
	return query.Updates(updateIn).Error
}
