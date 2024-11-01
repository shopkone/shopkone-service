package sPurchase

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/internal/module/product/purchase/sPurchase/sPurchaseItem"
	"shopkone-service/utility/handle"
)

// 获取采购单列表
func (s *sPurchase) List(in vo.PurchaseListReq) (out handle.PageRes[vo.PurchaseListRes], err error) {
	var list []mPurchase.Purchase
	query := s.orm.Model(&list).Where("shop_id =?", s.shopId)

	// 查找数量
	if err = query.Count(&out.Total).Error; err != nil {
		return out, err
	}

	// 分页
	query = query.Scopes(handle.Pagination(in.PageReq)).Order("created_at DESC")
	query = query.Select(
		"id",
		"destination_id",
		"purchase_number",
		"supplier_id",
		"estimated_arrival",
		"status",
		"adjust",
	)

	// 查询数据
	if err = query.Find(&list).Error; err != nil {
		return out, err
	}

	// 查询子项
	sItem := sPurchaseItem.NewPurchaseItem(s.orm, s.shopId)
	purchaseIds := slice.Map(list, func(_ int, item mPurchase.Purchase) uint { return item.ID })
	items, err := sItem.ListByPurchaseIds(purchaseIds)
	if err != nil {
		return out, err
	}

	// 组装数据
	out.List = slice.Map(list, func(_ int, item mPurchase.Purchase) vo.PurchaseListRes {
		currentItems := slice.Filter(items, func(index int, i mPurchase.PurchaseItem) bool {
			return i.PurchaseId == item.ID
		})
		purchasingTotal := 0
		var rejectTotal uint
		var receivedTotal uint
		for _, v := range currentItems {
			purchasingTotal += v.Purchasing
			rejectTotal += uint(v.Rejected)
			receivedTotal += uint(v.Received)
		}
		totalIn := GetTotalPriceIn{
			Items:   currentItems,
			Status:  item.Status,
			Adjusts: item.Adjust,
		}
		i := vo.PurchaseListRes{
			DestinationId: item.DestinationId,
			Status:        item.Status,
			SupplierId:    item.SupplierId,
			Id:            item.ID,
			OrderNumber:   item.PurchaseNumber,
			Purchasing:    purchasingTotal,
			Received:      receivedTotal,
			Rejected:      rejectTotal,
			Total:         s.GetTotalPrice(totalIn),
		}
		if item.EstimatedArrival != nil {
			i.EstimatedArrival = handle.ToUnix(item.EstimatedArrival)
		}
		return i
	})

	return out, nil
}
