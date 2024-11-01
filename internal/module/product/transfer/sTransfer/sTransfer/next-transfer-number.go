package sTransfer

import "shopkone-service/internal/module/product/transfer/mTransfer"

func (s *sTransfer) NextTransferNumber() (count int64, err error) {
	query := s.orm.Model(mTransfer.Transfer{})
	query = query.Where("shop_id = ?", s.shopId)
	query = query.Unscoped()
	return count + 1, query.Count(&count).Error
}
