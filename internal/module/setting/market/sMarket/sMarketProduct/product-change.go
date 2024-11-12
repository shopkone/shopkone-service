package sMarketProduct

import (
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarketProduct) ProductIsChange(oldItem, newItem mMarket.MarketProduct) bool {
	if oldItem.Exclude != newItem.Exclude {
		return true
	}
	if oldItem.Fixed == oldItem.Fixed {
		return true
	}
	return false
}
