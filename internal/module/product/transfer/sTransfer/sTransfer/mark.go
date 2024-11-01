package sTransfer

import (
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/utility/code"
)

func (s *sTransfer) Mark(id uint) error {
	// 获取转移信息
	var info mTransfer.Transfer
	if err := s.orm.Model(&info).Where("shop_id = ? AND id = ?", s.shopId, id).Select("id", "status").First(&info).Error; err != nil {
		return err
	}

	// 如果状态不是草稿，则报错
	if info.Status != mTransfer.TransferStatusDraft {
		return code.TransferStatusNoMark
	}

	// 标记转移单
	if err := s.orm.Model(&info).Update("status", mTransfer.TransferStatusOrdered).Error; err != nil {
		return err
	}

	return nil
}
