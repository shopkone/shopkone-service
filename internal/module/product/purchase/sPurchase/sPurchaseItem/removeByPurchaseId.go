package sPurchaseItem

import "shopkone-service/internal/module/product/purchase/mPurchase"

func (s *sPurchaseItem) RemoveByPurchaseId(purchaseId uint) error {
	query := s.orm.Model(&mPurchase.PurchaseItem{}).Where("purchase_id = ?", purchaseId)
	query = query.Where("rejected = 0 AND rejected = 0")
	query = query.Where("shop_id = ?", s.shopId)
	return query.Delete(&mPurchase.PurchaseItem{}).Error
}
