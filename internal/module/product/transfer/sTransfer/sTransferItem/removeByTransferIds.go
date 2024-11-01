package sTransferItem

import (
	"shopkone-service/internal/module/product/transfer/mTransfer"
)

func (s *sTransferItem) RemoveByTransferIds(ids []uint) error {
	query := s.orm.Model(&mTransfer.TransferItem{})
	query = query.Where("transfer_id IN ?", ids)
	query = query.Where("shop_id = ?", s.shopId)
	return query.Delete(&mTransfer.TransferItem{}).Error
}
