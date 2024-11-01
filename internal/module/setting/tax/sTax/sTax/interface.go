package sTax

import "gorm.io/gorm"

type sTax struct {
	orm    *gorm.DB
	shopId uint
}

func NewTax(orm *gorm.DB, shopId uint) *sTax {
	return &sTax{orm: orm, shopId: shopId}
}
