package mCustomer

import (
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/base/orm/mOrm"
)

// 主要联系方式类型
type MainConcatType string

const (
	MainConcatTypePhone MainConcatType = "phone"
	MainConcatTypeEmail MainConcatType = "email"
)

// 客户来源
type CustomerSource uint8

const (
	CustomerSourceManual   CustomerSource = 1 // 手动创建
	CustomerSourceRegister CustomerSource = 2 // 注册账号
)

// 客户生命周期动作
type CustomerEvent int32

const (
	CustomerEventRegister CustomerEvent = 1 // 注册账号
	CustomerEventCreate   CustomerEvent = 2 // 手动创建
)

// 客户
type Customer struct {
	mOrm.Model
	FirstName          string
	LastName           string
	Email              string
	PhoneCode          string
	PhoneIso2          string
	PhoneNumber        string
	MainConcat         MainConcatType
	Password           string
	Subscribe          bool
	CountryIso3        string
	Source             CustomerSource
	Address            []mAddress.Address  `gorm:"foreignKey:CustomerID"`
	CustomerFootprints []CustomerFootprint `gorm:"foreignKey:CustomerID"`
	TaxFree            bool
	Tags               []string `gorm:"serializer:json"`
	OrderIds           []uint   `gorm:"serializer:json"`
}

// 客户足迹
type CustomerFootprint struct {
	mOrm.Model
	Event      CustomerEvent
	CustomerID uint
}
