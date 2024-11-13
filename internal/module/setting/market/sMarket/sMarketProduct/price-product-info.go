package sMarketProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	exchang_rate "shopkone-service/internal/module/base/exchang-rate"
	"shopkone-service/internal/module/setting/market/mMarket"
)

type MarketProductIn struct {
	MarketID      uint
	StoreCurrency string
}

func (s *sMarketProduct) GetPrice(marketID uint, storeCurrency string) (out vo.MarketGetProductRes, err error) {
	// 获取价格调整
	price, err := s.PriceInfo(marketID)
	if err != nil {
		return out, err
	}
	// 获取商品列表
	marketProduct, err := s.ProductList(marketID)
	if err != nil {
		return out, err
	}
	out.AdjustPercent = price.AdjustPercent
	out.AdjustType = price.AdjustType
	out.CurrencyCode = price.CurrencyCode
	out.AdjustProducts = slice.Map(marketProduct, func(index int, item mMarket.MarketProduct) vo.MarketUpdateProductItem {
		i := vo.MarketUpdateProductItem{}
		i.Fixed = item.Fixed
		i.ProductID = item.ProductID
		i.Exclude = item.Exclude
		i.ID = item.ID
		return i
	})
	if out.AdjustProducts == nil {
		out.AdjustProducts = []vo.MarketUpdateProductItem{}
	}

	// 获取费率
	getRateIn := exchang_rate.GetRateIn{
		FromCode: storeCurrency,
		ToCode:   out.CurrencyCode,
	}
	exchange, err := exchang_rate.NewExchangeRate().GetRate(getRateIn)
	if err != nil {
		return out, err
	}
	if exchange.Support {
		out.ExchangeRate = exchange.Rate
		out.ExchangeRateTimeStamp = exchange.Timestamp
	}

	return out, err
}
