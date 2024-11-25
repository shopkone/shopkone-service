package sCustomerAddress

import (
	"gorm.io/gorm"
)

type sCustomerAddress struct {
	orm    *gorm.DB
	shopID uint
}

func NewCustomerAddress(shopID uint, orm *gorm.DB) *sCustomerAddress {
	return &sCustomerAddress{orm: orm, shopID: shopID}
}
