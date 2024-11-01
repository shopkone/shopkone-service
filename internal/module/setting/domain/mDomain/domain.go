package mDomain

import "shopkone-service/internal/module/base/orm/mOrm"

type DomainStatus int

const (
	DomainStatusConnectSuccess DomainStatus = iota + 1 // 连接成功
	DomainStatusConnectFailed                          // 连接失败
	DomainStatusDisconnect                             // 取消连接
)

type Domain struct {
	mOrm.Model
	Domain     string       `json:"domain" gorm:"index"`      // 域名
	Status     DomainStatus `json:"status" gorm:"index"`      // 状态
	IsShopkone bool         `json:"is_shopkone" gorm:"index"` // 是否 shopkone 自带域名
}
