package sBlockCountry

import "gorm.io/gorm"

type sBlockCountry struct {
	orm    *gorm.DB
	shopId uint
}

func NewBlockCountry(orm *gorm.DB, shopId uint) *sBlockCountry {
	return &sBlockCountry{
		orm:    orm,
		shopId: shopId,
	}
}
