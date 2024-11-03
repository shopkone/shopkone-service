package sCustomerTax

import "gorm.io/gorm"

type sCustomerTax struct {
	orm    *gorm.DB
	shopId uint
}

func NewCustomerTax(orm *gorm.DB, shopId uint) *sCustomerTax {
	return &sCustomerTax{orm: orm, shopId: shopId}
}
