package sOrder

import "gorm.io/gorm"

type sOrder struct {
	orm    *gorm.DB
	shopId uint
}

func NewOrder(db *gorm.DB, shopId uint) *sOrder {
	return &sOrder{orm: db, shopId: shopId}
}
