package sTransfer

import (
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/internal/module/product/transfer/sTransfer/sTransferItem"
	"shopkone-service/utility/code"
)

// 获取当前应该存在的状态
func (s *sTransfer) SetStatus(id uint, oldStatus mTransfer.TransferStatus) (err error) {
	if oldStatus == 0 {
		return code.TransferStatusMissing
	}

	// 获取子项
	items, err := sTransferItem.NewTransferItem(s.orm, s.shopId).ListByTransferIds([]uint{id})
	if err != nil {
		return err
	}

	// 获取子项的各个总数
	var (
		quantity,
		received,
		rejected uint
	)
	for _, item := range items {
		quantity += item.Quantity
		received += item.Received
		rejected += item.Rejected
	}

	status := oldStatus

	if received+rejected == 0 {
		return
	}

	// 如果加起来小于总数则是部分采购
	if received+rejected < quantity {
		status = mTransfer.TransferStatusPartialReceived
	}
	// 如果加起来等于总数则是全部采购
	if received+rejected >= quantity {
		status = mTransfer.TransferStatusReceived
	}

	if status != mTransfer.TransferStatusPartialReceived && status != mTransfer.TransferStatusReceived {
		return err
	}

	return s.orm.Model(mTransfer.Transfer{}).
		Where("shop_id = ? AND id = ?", s.shopId, id).
		Update("status", status).Error
}
