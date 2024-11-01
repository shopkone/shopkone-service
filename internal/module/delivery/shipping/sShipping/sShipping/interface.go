package sShipping

import "gorm.io/gorm"

type sShipping struct {
	orm    *gorm.DB
	shopId uint
}

func NewShipping(orm *gorm.DB, shopId uint) *sShipping {
	return &sShipping{orm: orm, shopId: shopId}
}
