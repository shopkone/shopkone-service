package sTransfer

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/internal/module/product/transfer/sTransfer/sTransferItem"
	"shopkone-service/utility/code"
)

func (s *sTransfer) Adjust(in vo.TransferAdjustReq) error {
	// 获取转移信息
	var info mTransfer.Transfer
	if err := s.orm.Model(&info).Where("shop_id = ? AND id = ?", s.shopId, in.Id).Select("id", "status").First(&info).Error; err != nil {
		return err
	}

	// 只有状态为0或者状态为Draft，则不允许调整
	if info.Status == 0 || info.Status == mTransfer.TransferStatusDraft {
		return code.TransferStatusNotAdjust
	}

	// 更新子项
	if err := sTransferItem.NewTransferItem(s.orm, s.shopId).Adjust(in); err != nil {
		return err
	}

	// 更新状态
	return s.SetStatus(in.Id, info.Status)
}
