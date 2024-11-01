package sPurchaseItem

import "gorm.io/gorm"

type sPurchaseItem struct {
	orm    *gorm.DB
	shopId uint
}

type IPurchaseItem interface {
}

func NewPurchaseItem(orm *gorm.DB, shopId uint) *sPurchaseItem {
	return &sPurchaseItem{orm: orm, shopId: shopId}
}
