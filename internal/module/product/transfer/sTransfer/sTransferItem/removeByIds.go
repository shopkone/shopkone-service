package sTransferItem

import "shopkone-service/internal/module/product/transfer/mTransfer"

func (s *sTransferItem) RemoveByIds(ids []uint) error {
	query := s.orm.Model(&mTransfer.TransferItem{}).Where("id IN (?)", ids)
	query = query.Where("shop_id = ?", s.shopId)
	return query.Delete(&mTransfer.TransferItem{}).Error
}
