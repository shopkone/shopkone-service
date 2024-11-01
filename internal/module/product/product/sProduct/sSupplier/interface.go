package sSupplier

import "gorm.io/gorm"

type sSupplier struct {
	orm    *gorm.DB
	shopId uint
}

func NewSupplier(orm *gorm.DB, shopId uint) *sSupplier {
	return &sSupplier{orm: orm, shopId: shopId}
}

type ISupplier interface {
}
