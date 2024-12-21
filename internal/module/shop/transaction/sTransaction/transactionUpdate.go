package sTransaction

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/shop/transaction/mTransaction"
)

func (s *sTransaction) UpdateTransaction(in vo.ShopUpdateTransactionReq) (err error) {
	transaction := mTransaction.Transaction{
		AutoFinishDay:               in.AutoFinishDay,
		IsAutoFinish:                in.IsAutoFinish,
		IsForceCheckProduct:         in.IsForceCheckProduct,
		OrderAutoCancel:             in.OrderAutoCancel,
		OrderAutoCancelCustomerHour: in.OrderAutoCancelCustomerHour,
		ReduceTime:                  in.ReduceTime,
		TargetType:                  in.TargetType,
	}
	return s.orm.
		Select(
			"auto_finish_day",
			"is_auto_finish",
			"is_force_check_product",
			"order_auto_cancel",
			"order_auto_cancel_customer_hour",
			"reduce_time",
			"target_type",
		).
		Updates(&transaction).Error
}
