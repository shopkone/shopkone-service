package sInStorePickup

import "gorm.io/gorm"

type sInStorePickup struct {
	orm    *gorm.DB
	shopId uint
}

func NewInStorePickup(orm *gorm.DB, shopId uint) *sInStorePickup {
	return &sInStorePickup{orm: orm, shopId: shopId}
}
