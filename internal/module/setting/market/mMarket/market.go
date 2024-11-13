package mMarket

import (
	"shopkone-service/internal/module/base/orm/mOrm"
)

type MarketStatus uint8

const (
	MarketStatusActive   MarketStatus = iota + 1 // 激活
	MarketStatusInactive                         // 禁用
)

type DomainType uint8

const (
	DomainTypeMain   DomainType = iota + 1 // 使用主域名
	DomainTypeSub                          // 使用子域名
	DomainTypeSuffix                       // 使用后缀
)

type Market struct {
	mOrm.Model
	Name   string       `gorm:"comment:名称"`
	Status MarketStatus `gorm:"comment:状态"`
	IsMain bool         `gorm:"index;comment:是否主市场"`
	// 域名
	DomainType   DomainType `gorm:"default:1"`
	DomainSuffix string     `gorm:"size:50"`
	SubDomainID  uint       `gorm:"index"`
	// 语言
	DefaultLanguageID uint   `gorm:"index;not null"`
	LanguageIds       []uint `gorm:"serializer:json"`
}
