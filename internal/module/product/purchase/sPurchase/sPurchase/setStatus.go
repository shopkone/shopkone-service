package sPurchase

import (
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/internal/module/product/purchase/sPurchase/sPurchaseItem"
	"shopkone-service/utility/code"
)

// 获取当前应该存在的状态
func (s *sPurchase) SetStatus(id uint, oldStatus mPurchase.PurchaseStatus) (err error) {
	if oldStatus == 0 {
		return code.PurchaseStatusMissing
	}

	// 如果状态为关闭状态不处理
	if oldStatus == mPurchase.PurchaseStatusClosed {
		return err
	}

	// 获取子项
	items, err := sPurchaseItem.NewPurchaseItem(s.orm, s.shopId).ListByPurchaseId(id)
	if err != nil {
		return err
	}

	// 获取子项的各个总数
	var (
		purchasing,
		received,
		rejected int
	)
	for _, item := range items {
		purchasing += item.Purchasing
		received += item.Received
		rejected += item.Rejected
	}

	status := oldStatus

	if received+rejected == 0 {
		return
	}

	// 如果加起来小于总数则是部分采购
	if received+rejected < purchasing {
		status = mPurchase.PurchaseStatusPartialReceived
	}
	// 如果加起来等于总数则是全部采购
	if received+rejected >= purchasing {
		status = mPurchase.PurchaseStatusReceived
	}

	if status != mPurchase.PurchaseStatusPartialReceived && status != mPurchase.PurchaseStatusReceived {
		return err
	}

	return s.orm.Model(mPurchase.Purchase{}).
		Where("shop_id = ? AND id = ?", s.shopId, id).
		Update("status", status).Error
}
