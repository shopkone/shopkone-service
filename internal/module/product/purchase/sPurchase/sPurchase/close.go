package sPurchase

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/product/purchase/mPurchase"
)

func (s *sPurchase) Close(id uint, isClose bool) error {
	// 获取采购单信息
	var info mPurchase.Purchase
	if err := s.orm.Model(&info).Where("shop_id = ? AND id = ?", s.shopId, id).
		Select("id", "status", "old_status").First(&info).Error; err != nil {
		return err
	}

	// 如果是草稿订单或者为0，则不允许设置
	if info.Status == mPurchase.PurchaseStatusDraft || info.Status == 0 {
		return nil
	}

	var status mPurchase.PurchaseStatus
	var oldStatus mPurchase.PurchaseStatus
	if isClose {
		status = mPurchase.PurchaseStatusClosed
		oldStatus = info.Status
	} else {
		status = info.OldStatus
	}

	if status == 0 {
		return nil
	}

	updateIn := g.Map{
		"status":     status,
		"old_status": oldStatus,
	}

	query := s.orm.Model(&mPurchase.Purchase{}).Where("shop_id = ? AND id = ?", s.shopId, id)
	return query.Updates(updateIn).Error
}
