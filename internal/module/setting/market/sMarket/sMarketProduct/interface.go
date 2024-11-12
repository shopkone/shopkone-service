package sMarketProduct

import (
	"gorm.io/gorm"
)

type sMarketProduct struct {
	orm    *gorm.DB
	shopId uint
}

func NewMarketProduct(orm *gorm.DB, shopId uint) *sMarketProduct {
	return &sMarketProduct{orm: orm, shopId: shopId}
}
