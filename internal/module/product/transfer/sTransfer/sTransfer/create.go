package sTransfer

import (
	"github.com/duke-git/lancet/v2/convertor"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/internal/module/product/transfer/sTransfer/sTransferItem"
	"shopkone-service/internal/module/setting/location/sLocation"
	"shopkone-service/utility/handle"
)

func (s *sTransfer) CreateTransfer(in vo.TransferCreateReq) (uint, error) {
	locationService := sLocation.NewLocation(s.orm, s.shopId)

	// 校验地点是否可用
	locationIds := []uint{in.OriginId, in.DestinationId}
	if err := locationService.IsAllActive(locationIds); err != nil {
		return 0, err
	}

	// 创建库存转移单
	transferNumber, err := s.NextTransferNumber()
	if err != nil {
		return 0, err
	}
	data := &mTransfer.Transfer{}
	data.ShopId = s.shopId
	data.OriginId = in.OriginId
	data.DestinationId = in.DestinationId
	data.TransferNumber = "#TR" + convertor.ToString(transferNumber)
	data.Status = mTransfer.TransferStatusDraft
	data.CarrierId = in.CarrierId
	data.DeliveryNumber = in.DeliveryNumber
	if in.EstimatedArrival != 0 {
		data.EstimatedArrival = handle.ParseTime(in.EstimatedArrival)
	}

	if err = s.orm.Create(data).Error; err != nil {
		return 0, err
	}

	// 创建转移单明细
	transferItemService := sTransferItem.NewTransferItem(s.orm, s.shopId)
	createIn := sTransferItem.CreateTransferItemsIn{
		Items:         in.Items,
		TransferId:    data.ID,
		OriginId:      in.OriginId,
		DestinationId: in.DestinationId,
	}
	if err = transferItemService.CreateTransferItems(createIn); err != nil {
		return 0, err
	}

	return data.ID, nil
}
