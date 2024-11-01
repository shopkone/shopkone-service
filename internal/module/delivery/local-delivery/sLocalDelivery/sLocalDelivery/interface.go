package sLocalDelivery

import "gorm.io/gorm"

type sLocalDelivery struct {
	orm    *gorm.DB
	shopId uint
}

func NewLocalDelivery(orm *gorm.DB, shopId uint) *sLocalDelivery {
	return &sLocalDelivery{orm: orm, shopId: shopId}
}
