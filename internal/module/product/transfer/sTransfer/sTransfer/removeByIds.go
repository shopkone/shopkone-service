package sTransfer

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/internal/module/product/transfer/sTransfer/sTransferItem"
	"shopkone-service/utility/code"
)

func (s *sTransfer) RemoveByIds(ids []uint) error {
	// 获取转移单
	var transfers []mTransfer.Transfer
	if err := s.orm.Model(&transfers).Where("shop_id = ? AND id IN ?", s.shopId, ids).
		Select("id", "status").Find(&transfers).Error; err != nil {
		return err
	}

	// 如果状态不是草稿单，则不允许删除
	isAllDraft := slice.Every(transfers, func(index int, item mTransfer.Transfer) bool {
		return item.Status == mTransfer.TransferStatusDraft
	})
	if !isAllDraft {
		return code.ErrTransferStatusNotDrafts
	}

	// 删除转移单
	query := s.orm.Model(&mTransfer.Transfer{})
	query = query.Where("id IN ?", ids)
	query = query.Where("shop_id = ?", s.shopId)
	if err := query.Delete(&mTransfer.Transfer{}).Error; err != nil {
		return err
	}

	// 删除子项
	return sTransferItem.NewTransferItem(s.orm, s.shopId).RemoveByTransferIds(ids)
}
