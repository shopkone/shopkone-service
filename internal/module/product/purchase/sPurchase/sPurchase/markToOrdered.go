package sPurchase

import (
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/utility/code"
)

func (s *sPurchase) MarkToOrdered(purchaseId uint) (err error) {
	// 获取采购单信息
	var purchase mPurchase.Purchase
	query := s.orm.Model(&purchase).Where("shop_id = ? AND id = ?", s.shopId, purchaseId)
	query = query.Select("status", "id")
	if err = query.First(&purchase).Error; err != nil {
		return err
	}

	// 只有采购订单为草稿状态才允许标记
	if purchase.Status != mPurchase.PurchaseStatusDraft {
		return code.PurchaseIsOrdered
	}

	// 更新采购单状态为已订购
	query = s.orm.Model(&purchase).Where("shop_id = ? AND id = ?", s.shopId, purchaseId)
	return query.Update("status", mPurchase.PurchaseStatusOrdered).Error
}
