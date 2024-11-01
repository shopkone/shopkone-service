package sPurchaseItem

import "shopkone-service/internal/module/product/purchase/mPurchase"

func (s *sPurchaseItem) ListByPurchaseId(purchaseId uint) (res []mPurchase.PurchaseItem, err error) {
	query := s.orm.Model(&res).Where("purchase_id = ?", purchaseId)
	query = query.Where("shop_id = ?", s.shopId)
	query = query.Omit("created_at", "updated_at", "deleted_at", "shop_id")
	return res, query.Find(&res).Error
}
