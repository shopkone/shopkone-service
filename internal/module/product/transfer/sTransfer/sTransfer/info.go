package sTransfer

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/internal/module/product/transfer/sTransfer/sTransferItem"
	"shopkone-service/utility/handle"
)

func (s *sTransfer) Info(in uint) (res vo.TransferInfoRes, err error) {
	// 查询数据
	var data mTransfer.Transfer
	query := s.orm.Model(mTransfer.Transfer{}).Where("shop_id =? AND id =?", s.shopId, in)
	query = query.Omit("shop_id", "created_at", "updated_at", "deleted_at")
	if err = query.First(&data).Error; err != nil {
		return res, err
	}

	// 获取子项
	items, err := sTransferItem.NewTransferItem(s.orm, s.shopId).ListByTransferIds([]uint{data.ID})
	if err != nil {
		return res, err
	}

	// 组装数据
	res = vo.TransferInfoRes{}
	res.Id = data.ID
	res.CarrierId = data.CarrierId
	res.OriginId = data.OriginId
	res.DestinationId = data.DestinationId
	res.Status = data.Status
	res.DeliveryNumber = data.DeliveryNumber
	if data.EstimatedArrival != nil {
		res.EstimatedArrival = handle.ToUnix(data.EstimatedArrival)
	}
	res.TransferNumber = data.TransferNumber
	res.Items = slice.Map(items, func(_ int, item mTransfer.TransferItem) vo.BaseTransferItem {
		i := vo.BaseTransferItem{}
		i.Id = item.ID
		i.Quantity = item.Quantity
		i.VariantId = item.VariantId
		i.Rejected = item.Rejected
		i.Received = item.Received
		return i
	})
	return res, nil
}
