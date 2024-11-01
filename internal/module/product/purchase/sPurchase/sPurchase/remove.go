package sPurchase

import (
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/internal/module/product/purchase/sPurchase/sPurchaseItem"
	"shopkone-service/utility/code"
)

func (s *sPurchase) Remove(purchaseId uint) (err error) {
	// 获取采购单信息
	var purchase mPurchase.Purchase
	query := s.orm.Model(&purchase).Where("shop_id = ? AND id = ?", s.shopId, purchaseId)
	query = query.Select("status", "id")
	if err = query.First(&purchase).Error; err != nil {
		return err
	}

	// 只有状态为草稿才可以删除
	if purchase.Status != mPurchase.PurchaseStatusDraft {
		return code.PurchaseCanNotRemove
	}

	// 删除采购单
	err = s.orm.Model(&mPurchase.Purchase{}).
		Where("shop_id = ? AND id = ?", s.shopId, purchaseId).
		Delete(&mPurchase.Purchase{}, purchaseId).Error
	if err != nil {
		return err
	}

	// 删除采购单子项目
	return sPurchaseItem.NewPurchaseItem(s.orm, s.shopId).RemoveByPurchaseId(purchaseId)
}
