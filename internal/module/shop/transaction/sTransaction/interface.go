package sTransaction

import "gorm.io/gorm"

type sTransaction struct {
	orm    *gorm.DB
	shopId uint
}

func NewTransaction(orm *gorm.DB, shopId uint) *sTransaction {
	return &sTransaction{orm: orm, shopId: shopId}
}
