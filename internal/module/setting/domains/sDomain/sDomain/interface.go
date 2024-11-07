package sDomain

import "gorm.io/gorm"

type sDomain struct {
	orm    *gorm.DB
	shopId uint
}

func NewDomain(orm *gorm.DB, shopId uint) *sDomain {
	return &sDomain{orm: orm, shopId: shopId}
}
