package sPurchase

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/sProduct/sSupplier"
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/internal/module/product/purchase/sPurchase/sPurchaseItem"
	"shopkone-service/internal/module/setting/location/sLocation"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

// 更新采购单
func (s *sPurchase) Update(in vo.PurchaseUpdateReq) (err error) {
	// 如果子项为0，则返回错误
	if len(in.PurchaseItems) == 0 {
		return code.PurchaseItemIsEmpty
	}

	// 如果输入了物流单号但是没有选择物流商，则返回错误
	if in.DeliveryNumber != "" && in.CarrierId == nil {
		return code.CarrierIdIsEmpty
	}

	// 调整项不能超过8项
	if len(in.Adjust) > 8 {
		return code.AdjustCountMax
	}

	// 校验目的是否启用
	if err = sLocation.NewLocation(s.orm, s.shopId).CheckExist(in.DestinationId); err != nil {
		return err
	}

	// 校验供应商
	if err = sSupplier.NewSupplier(s.orm, s.shopId).CheckExist(in.SupplierId); err != nil {
		return err
	}

	// 获取采购单信息
	var purchase mPurchase.Purchase
	if err = s.orm.Model(&purchase).Where("shop_id = ? AND id = ?", s.shopId, in.Id).Select("id", "status", "destination_id", "supplier_id").First(&purchase).Error; err != nil {
		return err
	}

	// 如果不是草稿单，则不允许修改地点或供应商
	if purchase.Status != mPurchase.PurchaseStatusDraft {
		if in.DestinationId != purchase.DestinationId {
			return code.DestinationIdNotUpdate
		}
		if in.SupplierId != purchase.SupplierId {
			return code.SupplierIdNotUpdate
		}
	}

	// 更新采购单
	data := mPurchase.Purchase{}
	data.CarrierId = in.CarrierId
	data.CurrencyCode = in.CurrencyCode
	data.DestinationId = in.DestinationId
	data.PaymentTerms = in.PaymentTerms
	data.Remarks = in.Remarks
	data.SupplierId = in.SupplierId
	data.Adjust = in.Adjust
	data.DeliveryNumber = in.DeliveryNumber
	if in.EstimatedArrival != 0 {
		data.EstimatedArrival = handle.ParseTime(in.EstimatedArrival)
	}
	if err = s.orm.Model(&data).Select(
		"carrier_id",
		"currency_code",
		"destination_id",
		"payment_terms",
		"remarks",
		"supplier_id",
		"adjust",
		"estimated_arrival",
		"delivery_number",
	).Where("shop_id = ? AND id = ?", s.shopId, in.Id).Updates(&data).Error; err != nil {
		return err
	}

	// 更新采购项
	updateIn := sPurchaseItem.PurchaseItemUpdateIn{
		Items:      in.PurchaseItems,
		PurchaseId: in.Id,
	}
	if err = sPurchaseItem.NewPurchaseItem(s.orm, s.shopId).Update(updateIn); err != nil {
		return err
	}

	// 更新状态
	if err = s.SetStatus(in.Id, purchase.Status); err != nil {
		return err
	}

	return err
}
