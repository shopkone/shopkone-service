package sPurchaseItem

import "shopkone-service/internal/module/product/purchase/mPurchase"

func (s *sPurchaseItem) ListByPurchaseIds(purchaseIds []uint) (res []mPurchase.PurchaseItem, err error) {
	query := s.orm.Model(&res).Where("purchase_id in ?", purchaseIds)
	query = query.Where("shop_id = ?", s.shopId)
	query = query.Select("id", "purchasing", "cost", "tax_rate", "purchase_id", "rejected", "received")
	return res, query.Find(&res).Error
}
