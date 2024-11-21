package sCustomer

import "gorm.io/gorm"

type sCustomer struct {
	orm    *gorm.DB
	shopId uint
}

func NewCustomer(orm *gorm.DB, shopId uint) *sCustomer {
	return &sCustomer{orm: orm, shopId: shopId}
}
