package sOrderDiscount

import "gorm.io/gorm"

type sOrderDiscount struct {
	db     *gorm.DB
	shopId uint
}

func NewOrderDiscount(db *gorm.DB, shopId uint) *sOrderDiscount {
	return &sOrderDiscount{db: db, shopId: shopId}
}
