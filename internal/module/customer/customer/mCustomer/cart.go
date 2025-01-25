package mCustomer

import "shopkone-service/internal/module/base/orm/mOrm"

type Cart struct {
	mOrm.Model
	OrderID uint // 订单id
}
