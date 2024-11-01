package sPurchaseItem

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/sProduct/sProduct"
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sPurchaseItem) Create(items []vo.BasePurchaseItem, purchaseId uint) error {
	data := slice.Map(items, func(index int, item vo.BasePurchaseItem) mPurchase.PurchaseItem {
		i := mPurchase.PurchaseItem{}
		i.ShopId = s.shopId
		i.PurchaseId = purchaseId
		i.VariantID = item.VariantID
		i.Cost = handle.RoundMoney(item.Cost)
		i.Purchasing = item.Purchasing
		i.TaxRate = item.TaxRate
		i.SKU = item.SKU
		return i
	})
	// 判断变体是否都跟踪库存
	variantIds := slice.Map(data, func(index int, item mPurchase.PurchaseItem) uint {
		return item.VariantID
	})
	is, err := sProduct.NewProduct(s.orm, s.shopId).IsAllTrack(variantIds)
	if err != nil {
		return err
	}
	if !is {
		return code.ProductNotTrackInventory
	}
	return s.orm.Create(&data).Error
}
