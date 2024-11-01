package sTransferItem

import "shopkone-service/internal/module/product/transfer/mTransfer"

func (s *sTransferItem) ListByTransferIds(transferIds []uint) ([]mTransfer.TransferItem, error) {
	var items []mTransfer.TransferItem
	query := s.orm.Model(&mTransfer.TransferItem{}).Where("transfer_id IN ?", transferIds)
	query = query.Where("shop_id = ?", s.shopId)
	query = query.Omit("created_at", "updated_at", "deleted_at", "shop_id")
	query = query.Order("id ASC")
	return items, query.Find(&items).Error
}
