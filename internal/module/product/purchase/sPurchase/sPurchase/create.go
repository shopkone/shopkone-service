package sPurchase

import (
	"github.com/duke-git/lancet/v2/convertor"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/sProduct/sSupplier"
	"shopkone-service/internal/module/product/purchase/mPurchase"
	"shopkone-service/internal/module/product/purchase/sPurchase/sPurchaseItem"
	"shopkone-service/internal/module/setting/location/sLocation"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

// 创建采购单
func (s *sPurchase) Create(in vo.PurchaseCreateReq) (id uint, err error) {
	// 如果子项为0，则返回错误
	if len(in.PurchaseItems) == 0 {
		return 0, code.PurchaseItemIsEmpty
	}

	// 如果输入了物流单号但是没有选择物流商，则返回错误
	if in.DeliveryNumber != "" && in.CarrierId == nil {
		return 0, code.CarrierIdIsEmpty
	}

	// 调整项不能超过8项
	if len(in.Adjust) > 8 {
		return 0, code.AdjustCountMax
	}

	// 校验目的是否启用
	if err = sLocation.NewLocation(s.orm, s.shopId).CheckExist(in.DestinationId); err != nil {
		return 0, err
	}

	// 校验供应商
	if err = sSupplier.NewSupplier(s.orm, s.shopId).CheckExist(in.SupplierId); err != nil {
		return 0, err
	}

	// 获取采购单数量，用于生成采购单号
	var count int64
	if err = s.orm.Model(&mPurchase.Purchase{}).Unscoped().Where("shop_id =?", s.shopId).Count(&count).Error; err != nil {
		return 0, err
	}

	// 创建采购单
	data := mPurchase.Purchase{}
	data.Status = mPurchase.PurchaseStatusDraft
	data.PurchaseNumber = "#PO" + convertor.ToString(count+1)
	data.ShopId = s.shopId
	data.CarrierId = in.CarrierId
	data.CurrencyCode = in.CurrencyCode
	data.DestinationId = in.DestinationId
	data.PaymentTerms = in.PaymentTerms
	data.Remarks = in.Remarks
	data.SupplierId = in.SupplierId
	data.DeliveryNumber = in.DeliveryNumber
	data.Adjust = in.Adjust
	if in.EstimatedArrival != 0 {
		data.EstimatedArrival = handle.ParseTime(in.EstimatedArrival)
	}
	if err = s.orm.Model(&data).Create(&data).Error; err != nil {
		return 0, err
	}

	// 创建采购单子项
	if err = sPurchaseItem.NewPurchaseItem(s.orm, s.shopId).Create(in.PurchaseItems, data.ID); err != nil {
		return 0, err
	}

	return data.ID, err
}
