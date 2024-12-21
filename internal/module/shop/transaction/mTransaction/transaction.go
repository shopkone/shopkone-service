package mTransaction

import "shopkone-service/internal/module/base/orm/mOrm"

type TransactionTargetType uint8

const (
	TransactionTargetTypeAll   TransactionTargetType = 1 // 所有客户
	TransactionTargetTypeLogin TransactionTargetType = 2 // 仅登录客户
)

type TransactionReduceTime uint8

const (
	TransactionReduceTimePay   TransactionReduceTime = 1 // 支付时扣减库存
	TransactionReduceTimeOrder TransactionReduceTime = 2 // 下单时扣减库存
)

type OrderAutoCancelType int32

const (
	OrderAutoCancel30Min   OrderAutoCancelType = 30
	OrderAutoCancel24Hour  OrderAutoCancelType = 1440
	OrderAutoCancel48Hour  OrderAutoCancelType = 2880
	OrderAutoCancelUnlimit OrderAutoCancelType = 0 // 不限时长
	OrderAutoCancelCustom  OrderAutoCancelType = -1
)

type Transaction struct {
	mOrm.Model
	TargetType                  TransactionTargetType // 目标客户
	ReduceTime                  TransactionReduceTime // 库存扣减时机
	IsForceCheckProduct         bool                  // 是否强制检测商品库存
	IsAutoFinish                bool                  // 是否开启自动收货
	AutoFinishDay               uint                  // 自动收货时间
	OrderAutoCancel             OrderAutoCancelType   // 订单自动取消
	OrderAutoCancelCustomerHour uint                  // 订单自动取消自动逸时间
}
