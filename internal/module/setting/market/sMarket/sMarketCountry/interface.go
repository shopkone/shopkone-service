package sMarketCountry

import "gorm.io/gorm"

type sMarketCountry struct {
	orm    *gorm.DB
	shopId uint
}

func NewMarketCountry(orm *gorm.DB, shopId uint) *sMarketCountry {
	return &sMarketCountry{orm: orm, shopId: shopId}
}
