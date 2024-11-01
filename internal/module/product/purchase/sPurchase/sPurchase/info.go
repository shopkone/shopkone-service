package sPurchase

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/internal/module/product/purchase/sPurchase/sPurchaseItem"
	"shopkone-service/utility/handle"
)

func (s *sPurchase) Info(id uint) (res vo.PurchaseInfoRes, err error) {
	// 获取采购单
	var purchase mPurchase.Purchase
	query := s.orm.Model(&purchase).Where("id = ?", id)
	query = query.Where("shop_id = ?", s.shopId)
	query = query.Omit("shop_id", "created_at", "deleted_at", "updated_at")
	if err = query.First(&purchase).Error; err != nil {
		return res, err
	}

	// 获取采购单子项
	sItem := sPurchaseItem.NewPurchaseItem(s.orm, s.shopId)
	items, err := sItem.ListByPurchaseId(id)
	if err != nil {
		return res, err
	}

	// 组装数据
	res = vo.PurchaseInfoRes{
		PurchaseCreateReq: vo.PurchaseCreateReq{
			SupplierId:     purchase.SupplierId,
			DestinationId:  purchase.DestinationId,
			CarrierId:      purchase.CarrierId,
			DeliveryNumber: purchase.DeliveryNumber,
			CurrencyCode:   purchase.CurrencyCode,
			Remarks:        purchase.Remarks,
			PaymentTerms:   purchase.PaymentTerms,
			Adjust:         purchase.Adjust,
		},
		Id:          purchase.ID,
		Status:      purchase.Status,
		OrderNumber: purchase.PurchaseNumber,
	}
	res.PurchaseItems = slice.Map(items, func(index int, item mPurchase.PurchaseItem) vo.BasePurchaseItem {
		i := vo.BasePurchaseItem{
			Id:         item.ID,
			VariantID:  item.VariantID,
			Purchasing: item.Purchasing,
			Cost:       item.Cost,
			SKU:        item.SKU,
			TaxRate:    item.TaxRate,
			Rejected:   item.Rejected,
			Received:   item.Received,
			Total:      sItem.CalculateTotal(item, res.Status),
			IsActive:   item.Active,
		}
		res.Received += item.Received
		res.Rejected += item.Rejected
		res.Purchasing += item.Purchasing
		return i
	})
	if purchase.EstimatedArrival != nil {
		res.EstimatedArrival = handle.ToUnix(purchase.EstimatedArrival)
	}
	return res, nil
}
