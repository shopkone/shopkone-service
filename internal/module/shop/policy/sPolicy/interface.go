package sPolicy

import "gorm.io/gorm"

type sPolicy struct {
	orm    *gorm.DB
	shopId uint
}

func NewPolicy(orm *gorm.DB, shopId uint) *sPolicy {
	return &sPolicy{orm: orm, shopId: shopId}
}
