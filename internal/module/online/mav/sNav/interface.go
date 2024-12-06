package sNav

import "gorm.io/gorm"

type sNav struct {
	orm    *gorm.DB
	shopId uint
}

func NewNav(orm *gorm.DB, shopId uint) *sNav {
	return &sNav{orm: orm, shopId: shopId}
}
