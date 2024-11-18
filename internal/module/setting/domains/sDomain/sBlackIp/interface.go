package sBlackIp

import "gorm.io/gorm"

type sBlockIp struct {
	orm    *gorm.DB
	shopId uint
}
