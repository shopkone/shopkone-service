package mLocation

import "shopkone-service/internal/module/base/orm/mOrm"

type LocationStatus int

const (
	LocationStatusOpen  LocationStatus = iota + 1 // 开启
	LocationStatusClose                           // 关闭
)

type Location struct {
	mOrm.Model
	AddressId          uint
	Active             bool
	Name               string
	IsDefault          bool
	FulfillmentDetails bool // 是否开启物流
	OrderNum           uint // 排序
}
