package sMarketProduct

import "shopkone-service/internal/module/setting/market/mMarket"

type MarketPriceUpdateIn struct {
	MarketID      uint
	CurrencyCode  string
	AdjustPercent float64
	AdjustType    mMarket.PriceAdjustmentType
	IsMain        bool
	StoreCurrency string
}

func (s *sMarketProduct) PriceUpdate(in MarketPriceUpdateIn) (err error) {
	if in.IsMain {
		// 如果是主市场，percent只能为0
		in.AdjustPercent = 0
		// 强制使用主货币
		in.CurrencyCode = in.StoreCurrency
	}
	updateIn := mMarket.MarketPrice{}
	updateIn.AdjustPercent = in.AdjustPercent
	updateIn.AdjustType = in.AdjustType
	updateIn.CurrencyCode = in.CurrencyCode
	query := s.orm.Model(&updateIn).Where("market_id = ?", in.MarketID)
	query = query.Select("adjust_percent", "adjust_type", "currency_code")
	return query.Updates(updateIn).Error
}
