package mDomains

import "shopkone-service/internal/module/base/orm/mOrm"

type BlackIpType uint8

const (
	BlackIpTypeBlack BlackIpType = iota + 1
	BlackIpTypeWhite
)

type DomainBlackIp struct {
	mOrm.Model
	Ip         string `gorm:"size:100"`
	ShowInList bool   `gorm:"default:true"`
	Type       BlackIpType
}
