package sShippingZone

import (
	"gorm.io/gorm"
)

type sShippingZone struct {
	shopId uint
	orm    *gorm.DB
}

func NewShippingZone(shopId uint, orm *gorm.DB) *sShippingZone {
	return &sShippingZone{shopId: shopId, orm: orm}
}
