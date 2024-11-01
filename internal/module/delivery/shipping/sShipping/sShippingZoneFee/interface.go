package sShippingZoneFee

import "gorm.io/gorm"

type sShippingZoneFee struct {
	orm    *gorm.DB
	shopId uint
}

func NewShippingZoneFee(orm *gorm.DB, shopId uint) *sShippingZoneFee {
	return &sShippingZoneFee{orm: orm, shopId: shopId}
}
