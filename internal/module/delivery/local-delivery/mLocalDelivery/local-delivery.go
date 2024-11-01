package mLocalDelivery

import (
	"shopkone-service/internal/module/base/orm/mOrm"
)

type LocalDeliveryStatus uint8

const (
	LocalDeliveryStatusOpen  LocalDeliveryStatus = 1 // 配送开启
	LocalDeliveryStatusClose LocalDeliveryStatus = 2 // 配送关闭
)

// LocalDelivery 本地配送
type LocalDelivery struct {
	mOrm.Model
	Status     LocalDeliveryStatus `json:"status"`
	LocationId uint                `gorm:"index;not null"`
}

// LocalDeliveryArea 配送区域
type LocalDeliveryArea struct {
	mOrm.Model
	Name            string `json:"name"`        // 区域名称
	PostalCode      string `json:"postal_code"` // 邮编
	Note            string `json:"note"`        // 备注
	LocalDeliveryID uint   `json:"local_delivery_id" gorm:"index;not null"`
}

// LocalDeliveryFee 配送配送
type LocalDeliveryFee struct {
	mOrm.Model
	Condition           float64 `json:"condition"` // 条件
	Fee                 float64 `json:"fee"`       // 配送费用
	LocalDeliveryAreaID uint    `json:"local_delivery_area_id" gorm:"index;not null"`
}
