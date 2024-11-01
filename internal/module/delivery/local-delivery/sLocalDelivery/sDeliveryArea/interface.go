package sDeliveryArea

import "gorm.io/gorm"

type sDeliveryArea struct {
	orm    *gorm.DB
	shopId uint
}

func NewDeliveryArea(orm *gorm.DB, shopId uint) *sDeliveryArea {
	return &sDeliveryArea{orm: orm, shopId: shopId}
}
