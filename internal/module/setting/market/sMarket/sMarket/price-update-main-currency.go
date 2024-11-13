package sMarket

import (
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/internal/module/setting/market/sMarket/sMarketProduct"
)

// 更新主要市场货币
func (s *sMarket) PriceUpdateMainCurrency(currency string) (err error) {
	var mainMarket mMarket.Market
	if err = s.orm.Model(&mMarket.Market{}).
		Where("shop_id = ? and is_main = ?", s.shopId, true).
		Select("id").
		First(&mainMarket).Error; err != nil {
		return err
	}
	priceUpdateIn := sMarketProduct.MarketPriceUpdateIn{
		MarketID:      mainMarket.ID,
		IsMain:        true,
		StoreCurrency: currency,
	}
	return sMarketProduct.NewMarketProduct(s.orm, s.shopId).PriceUpdate(priceUpdateIn)
}
