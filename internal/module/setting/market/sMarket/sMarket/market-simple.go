package sMarket

import "shopkone-service/internal/module/setting/market/mMarket"

type MarketSimpleOut struct {
	IsMain bool
}

func (s *sMarket) MarketSimple(marketId uint) (out MarketSimpleOut, err error) {
	var market mMarket.Market
	if err = s.orm.Model(&market).Where("id = ?", marketId).
		Select("is_main").
		First(&market).Error; err != nil {
		return out, err
	}
	out.IsMain = market.IsMain
	return out, nil
}
