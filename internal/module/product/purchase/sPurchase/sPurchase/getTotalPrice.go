package sPurchase

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/internal/module/product/purchase/sPurchase/sPurchaseItem"
)

type GetTotalPriceIn struct {
	Items   []mPurchase.PurchaseItem
	Status  mPurchase.PurchaseStatus
	Adjusts []mPurchase.PurchaseAdjustItem
}

func (s *sPurchase) GetTotalPrice(in GetTotalPriceIn) float64 {
	if in.Status == 0 || in.Status == mPurchase.PurchaseStatusDraft {
		return 0
	}
	sItem := sPurchaseItem.NewPurchaseItem(s.orm, s.shopId)
	var total float64
	slice.ForEach(in.Items, func(index int, item mPurchase.PurchaseItem) {
		total += sItem.CalculateTotal(item, in.Status)
	})
	slice.ForEach(in.Adjusts, func(index int, item mPurchase.PurchaseAdjustItem) {
		total += item.Price
	})
	return total
}
