package sMarketProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sMarketProduct) ProductUpdate(in vo.MarketUpdateProductReq, isMainMarket bool) error {
	// 获取价格调整
	marketPrice, err := s.PriceInfo(in.MarketID)
	if err != nil {
		return err
	}
	// 如果更改了货币
	if marketPrice.CurrencyCode != in.CurrencyCode {
		// 如果是主市场，则不允许再这里变更
		if isMainMarket {
			return code.MarketPriceUpdateInBaseCurrency
		}
		// 清空固定价格
		in.AdjustProducts = slice.Map(
			in.AdjustProducts,
			func(index int, item vo.MarketUpdateProductItem) vo.MarketUpdateProductItem {
				item.Fixed = nil
				return item
			})
	}
	// 更新价格调整
	priceUpdateIn := MarketPriceUpdateIn{
		MarketID:      in.MarketID,
		CurrencyCode:  in.CurrencyCode,
		AdjustPercent: in.AdjustPercent,
		AdjustType:    in.AdjustType,
	}
	if err = s.PriceUpdate(priceUpdateIn); err != nil {
		return err
	}
	// 获取旧的商品
	oldProducts, err := s.ProductList(in.MarketID)
	// 组装新的商品
	newProducts := slice.Map(in.AdjustProducts, func(index int, item vo.MarketUpdateProductItem) mMarket.MarketProduct {
		i := mMarket.MarketProduct{}
		i.ID = item.ID
		i.Exclude = item.Exclude
		i.Fixed = item.Fixed
		i.MarketID = in.MarketID
		i.ProductID = item.ProductID
		i.ShopId = s.shopId
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
