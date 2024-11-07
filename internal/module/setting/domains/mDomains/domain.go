package mDomains

import "shopkone-service/internal/module/base/orm/mOrm"

type DomainStatus uint8

const (
	DomainStatusConnectPre     DomainStatus = iota + 1 // 预连接
	DomainStatusConnectSuccess                         // 连接成功
	DomainStatusConnectFailed                          // 连接失败
	DomainStatusDisconnect                             // 连接断开
)

type Domain struct {
	mOrm.Model
	IsMain     bool         `gorm:"index"`
	Domain     string       `gorm:"size:500"`
	Status     DomainStatus `gorm:"default:0"`
	Ip         string       `gorm:"index"`
	IsShopKone bool         `gorm:"false"`
}
