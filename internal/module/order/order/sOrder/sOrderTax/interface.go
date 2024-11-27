package sOrderTax

import "gorm.io/gorm"

type sOrderTax struct {
	orm    *gorm.DB
	shopId uint
}

func NewOrderTax(db *gorm.DB, shopId uint) *sOrderTax {
	return &sOrderTax{
		orm:    db,
		shopId: shopId,
	}
}
