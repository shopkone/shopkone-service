package sTransferItem

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/transfer/mTransfer"
)

type CreateTransferItemsIn struct {
	Items         []vo.BaseTransferItem
	TransferId    uint
	OriginId      uint
	DestinationId uint
}

func (s *sTransferItem) CreateTransferItems(in CreateTransferItemsIn) error {
	// 创建项
	data := slice.Map(in.Items, func(_ int, item vo.BaseTransferItem) mTransfer.TransferItem {
		i := mTransfer.TransferItem{}
		i.ShopId = s.shopId
		i.TransferId = in.TransferId
		i.Quantity = item.Quantity
		i.VariantId = item.VariantId
		return i
	})
	if err := s.orm.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
