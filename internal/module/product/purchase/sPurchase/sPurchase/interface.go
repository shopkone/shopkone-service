package sPurchase

import "gorm.io/gorm"

type sPurchase struct {
	orm    *gorm.DB
	shopId uint
}

type IPurchase interface {
}

func NewPurchase(orm *gorm.DB, shopId uint) *sPurchase {
	return &sPurchase{orm: orm, shopId: shopId}
}
