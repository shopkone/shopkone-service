package sTransaction

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/shop/transaction/mTransaction"
)

func (s *sTransaction) Info() (res vo.ShopTransactionInfoRes, err error) {
	data := mTransaction.Transaction{}
	if err = s.orm.Where("shop_id = ?", s.shopId).
		Omit("shop_id", "created_at", "updated_at", "deleted_at").
		First(&data).Error; err != nil {
		return res, err
	}
	res.AutoFinishDay = data.AutoFinishDay
	res.IsAutoFinish = data.IsAutoFinish
	res.IsForceCheckProduct = data.IsForceCheckProduct
	res.OrderAutoCancel = data.OrderAutoCancel
	res.ReduceTime = data.ReduceTime
	res.TargetType = data.TargetType
	res.OrderAutoCancelCustomerHour = data.OrderAutoCancelCustomerHour
	return res, err
}
