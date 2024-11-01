package sTransfer

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/transfer/mTransfer"
	"shopkone-service/internal/module/product/transfer/sTransfer/sTransferItem"
	"shopkone-service/utility/handle"
)

func (s *sTransfer) List(in vo.TransferListReq) (res handle.PageRes[vo.TransferListRes], err error) {
	query := s.orm.Model(mTransfer.Transfer{}).Where("shop_id = ?", s.shopId)

	// 查找数量
	if err = query.Count(&res.Total).Error; err != nil {
		return res, err
	}

	// 分页
	query = query.Scopes(handle.Pagination(in.PageReq)).Order("created_at DESC")

	// 查询数据
	var list []mTransfer.Transfer
	query = query.Omit("created_at", "updated_at", "deleted_at", "shop_id")
	if err = query.Find(&list).Error; err != nil {
		return res, err
	}

	// 获取子项
	ids := slice.Map(list, func(_ int, item mTransfer.Transfer) uint { return item.ID })
	items, err := sTransferItem.NewTransferItem(s.orm, s.shopId).ListByTransferIds(ids)
	if err != nil {
		return res, err
	}

	// 转换数据
	res.List = slice.Map(list, func(_ int, item mTransfer.Transfer) vo.TransferListRes {
		i := vo.TransferListRes{}
		i.Id = item.ID
		i.Status = item.Status
		i.DestinationId = item.DestinationId
		i.OriginId = item.OriginId
		i.TransferNumber = item.TransferNumber
		if item.EstimatedArrival != nil {
			i.EstimatedArrival = handle.ToUnix(item.EstimatedArrival)
		}
		currentItems := slice.Filter(items, func(_ int, ii mTransfer.TransferItem) bool {
			return ii.TransferId == item.ID
		})
		slice.ForEach(currentItems, func(_ int, ii mTransfer.TransferItem) {
			i.Quantity += ii.Quantity
			i.Rejected += ii.Rejected
			i.Received += ii.Received
		})
		return i
	})
	return res, nil
}
