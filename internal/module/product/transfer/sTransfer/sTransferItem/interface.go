package sTransferItem

import "gorm.io/gorm"

type sTransferItem struct {
	shopId uint
	orm    *gorm.DB
}

func NewTransferItem(orm *gorm.DB, shopId uint) *sTransferItem {
	return &sTransferItem{shopId: shopId, orm: orm}
}
