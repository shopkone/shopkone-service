package sOrderShipping

import "gorm.io/gorm"

type sOrderShipping struct {
	orm    *gorm.DB
	shopId uint
}

func NewOrderShipping(db *gorm.DB, shopId uint) *sOrderShipping {
	return &sOrderShipping{
		orm:    db,
		shopId: shopId,
	}
}
