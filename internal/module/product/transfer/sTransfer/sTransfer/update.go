package sTransfer

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/internal/module/product/transfer/sTransfer/sTransferItem"
	"shopkone-service/utility/handle"
)

func (s *sTransfer) Update(in vo.TransferUpdateReq) error {
	// 获取转移信息
	var info mTransfer.Transfer
	if err := s.orm.Model(&info).Where("shop_id = ? AND id = ?", s.shopId, in.Id).Select("id", "status").First(&info).Error; err != nil {
		return err
	}

	// 仅能修改部分：不能修改商品和发货地以及目的地
	isUpdatePart := info.Status == mTransfer.TransferStatusPartialReceived || info.Status == mTransfer.TransferStatusReceived

	// 更新转移单
	info.DestinationId = in.DestinationId
	info.OriginId = in.OriginId
	info.CarrierId = in.CarrierId
	info.DeliveryNumber = in.DeliveryNumber
	if in.EstimatedArrival != 0 {
		info.EstimatedArrival = handle.ParseTime(in.EstimatedArrival)
	}

	query := s.orm.Model(&mTransfer.Transfer{})
	query = query.Where("shop_id =? AND id =?", s.shopId, in.Id)
	if isUpdatePart {
		query = query.Select("carrier_id", "delivery_number", "estimated_arrival")
	} else {
		query = query.Select("destination_id", "origin_id", "carrier_id", "delivery_number", "estimated_arrival")
	}
	if err := query.Updates(&info).Error; err != nil {
		return err
	}

	if isUpdatePart {
		return nil
	}

	// 更新子项
	updateItemsIn := sTransferItem.ItemUpdateIn{
		TransferId:    info.ID,
		Items:         in.Items,
		OriginId:      info.OriginId,
		DestinationId: info.DestinationId,
	}
	return sTransferItem.NewTransferItem(s.orm, s.shopId).Update(updateItemsIn)
}
