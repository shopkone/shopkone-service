package sMarket

import "gorm.io/gorm"

type sMarket struct {
	orm    *gorm.DB
	shopId uint
}

func NewMarket(orm *gorm.DB, shopId uint) *sMarket {
	return &sMarket{orm, shopId}
}
