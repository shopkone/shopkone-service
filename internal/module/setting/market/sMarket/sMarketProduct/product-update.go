package sMarketProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/handle"
)

func (s *sMarketProduct) ProductUpdate(in vo.MarketUpdateProductReq) error {
	// 更新定价调整数据
	market := mMarket.Market{}
	market.AdjustPercent = in.AdjustPercent
	market.AdjustType = in.AdjustType
	market.CurrencyCode = in.CurrencyCode
	if err := s.orm.Where("id = ?", in.MarketID).
		Select("adjust_percent", "adjust_type", "currency_code").
		Updates(&market).Error; err != nil {
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
