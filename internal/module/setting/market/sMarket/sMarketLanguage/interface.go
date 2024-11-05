package sMarketLanguage

import "gorm.io/gorm"

type sMarketLanguage struct {
	orm    *gorm.DB
	shopId uint
}

func NewMarketLanguage(orm *gorm.DB, shopId uint) *sMarketLanguage {
	return &sMarketLanguage{orm: orm, shopId: shopId}
}
