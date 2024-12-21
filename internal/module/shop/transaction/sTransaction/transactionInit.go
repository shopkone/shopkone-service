package sTransaction

import "shopkone-service/internal/module/shop/transaction/mTransaction"

func (s *sTransaction) Init() error {
	data := mTransaction.Transaction{
		AutoFinishDay:       0,
		IsAutoFinish:        false,                                 // 不开启自动收货
		IsForceCheckProduct: false,                                 // 不开启强制检测商品
		OrderAutoCancel:     mTransaction.OrderAutoCancel30Min,     // 30 分钟后订单自动失效
		ReduceTime:          mTransaction.TransactionReduceTimePay, // 支付时扣减库存
		TargetType:          mTransaction.TransactionTargetTypeAll, // 全部客户均可下单
	}
	data.ShopId = s.shopId
	return s.orm.Create(&data).Error
}
