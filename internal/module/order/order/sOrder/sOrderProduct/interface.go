package sOrderProduct

import "gorm.io/gorm"

type sOrderProduct struct {
	orm    *gorm.DB
	shopId uint
}

func NewOrderVariant(db *gorm.DB, shopId uint) *sOrderProduct {
	return &sOrderProduct{orm: db, shopId: shopId}
}
