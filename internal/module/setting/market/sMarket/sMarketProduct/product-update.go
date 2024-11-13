package sMarketProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/handle"
)

type ProductUpdateIn struct {
	Req           *vo.MarketUpdateProductReq
	StoreCurrency string
	IsMain        bool
}

func (s *sMarketProduct) ProductUpdate(in ProductUpdateIn) error {
	// 获取价格调整
	marketPrice, err := s.PriceInfo(in.Req.MarketID)
	if err != nil {
		return err
	}
	// 是否调整了货币
	isCurrencyChange := marketPrice.CurrencyCode != in.Req.CurrencyCode
	// 更新价格调整
	priceUpdateIn := MarketPriceUpdateIn{
		MarketID:      in.Req.MarketID,
		CurrencyCode:  in.Req.CurrencyCode,
		AdjustPercent: in.Req.AdjustPercent,
		AdjustType:    in.Req.AdjustType,
		StoreCurrency: in.StoreCurrency,
		IsMain:        in.IsMain,
	}
	if err = s.PriceUpdate(priceUpdateIn); err != nil {
		return err
	}

	// 获取旧的商品
	oldProducts, err := s.ProductList(in.Req.MarketID)
	// 组装新的商品
	newProducts := slice.Map(in.Req.AdjustProducts, func(index int, item vo.MarketUpdateProductItem) mMarket.MarketProduct {
		i := mMarket.MarketProduct{}
		i.ID = item.ID
		i.Exclude = item.Exclude
		i.Fixed = item.Fixed
		i.MarketID = in.Req.MarketID
		i.ProductID = item.ProductID
		i.ShopId = s.shopId
		// 如果是主市场或者更改了货币，则不允许有fixed
		if in.IsMain || isCurrencyChange {
			i.Fixed = nil
		}
		return i
	})
	insert, update, remove, err := handle.DiffUpdate(newProducts, oldProducts)
	if err != nil {
		return err
	}
	removeIds := slice.Map(remove, func(index int, item mMarket.MarketProduct) uint {
		return item.ID
	})
	if err = s.ProductRemove(removeIds); err != nil {
		return err
	}
	if err = s.ProductUpdateBatch(update, oldProducts); err != nil {
		return err
	}
	return s.ProductCreate(insert)
}
