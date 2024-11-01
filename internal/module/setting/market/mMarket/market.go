package mMarket

import "shopkone-service/internal/module/base/orm/mOrm"

type MarketStatus uint8

const (
	MarketStatusActive   MarketStatus = iota + 1 // 激活
	MarketStatusInactive                         // 禁用
)

type Market struct {
	mOrm.Model
	Name         string       `gorm:"comment:名称"`
	CountryCodes []string     `gorm:"comment:国家代码;serializer:JSON"`
	Status       MarketStatus `gorm:"comment:状态"`
	IsMain       bool         `gorm:"index;comment:是否主市场"`
}
