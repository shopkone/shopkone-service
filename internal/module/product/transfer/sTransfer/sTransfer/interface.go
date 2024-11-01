package sTransfer

import "gorm.io/gorm"

type sTransfer struct {
	orm    *gorm.DB
	shopId uint
}

func NewTransfer(orm *gorm.DB, shopId uint) *sTransfer {
	return &sTransfer{orm: orm, shopId: shopId}
}
