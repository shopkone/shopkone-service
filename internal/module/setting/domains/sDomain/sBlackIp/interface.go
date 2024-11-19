package sBlackIp

import "gorm.io/gorm"

type sBlackIp struct {
	orm    *gorm.DB
	shopId uint
}

func NewBlackIp(orm *gorm.DB, shopId uint) *sBlackIp {
	return &sBlackIp{
		orm:    orm,
		shopId: shopId,
	}
}
