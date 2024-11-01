package sPurchaseItem

import "shopkone-service/internal/module/product/purchase/mPurchase"

func (s *sPurchaseItem) IsChange(oldItem, newItem mPurchase.PurchaseItem) bool {
	if oldItem.Cost != newItem.Cost {
		return true
	}
	if oldItem.Purchasing != newItem.Purchasing {
		return true
	}
	if oldItem.SKU != newItem.SKU {
		return true
	}
	if oldItem.TaxRate != newItem.TaxRate {
		return true
	}
	return false
}
