package mCustomer

import (
	"shopkone-service/internal/module/base/orm/mOrm"
	"time"
)

type GenderType uint8

const (
	GenderTypeUnknown GenderType = 0
	GenderTypeMale    GenderType = 1
	GenderTypeFemale  GenderType = 2
	GenderTypeOther   GenderType = 3
)

// 客户
type Customer struct {
	mOrm.Model
	FirstName     string   `gorm:"size:200"`
	LastName      string   `gorm:"size:200"`
	Email         string   `gorm:"size:200"`
	Note          string   `gorm:"size:500"`
	Phone         string   `gorm:"size:100"`
	Tags          []string `gorm:"serializer:json"`
	Language      string   `gorm:"size:10"`
	Gender        GenderType
	Birthday      *time.Time
	TaxFreeAll    bool   `gorm:"default:false"`
	TaxFreeActive bool   `gorm:"default:false"`
	Password      string `gorm:"size:255"`
}

type CustomerAddress struct {
	mOrm.Model
	CustomerID uint `gorm:"index"`
	AddressID  uint `gorm:"index"`
	IsDefault  bool
}

type CustomerNoTaxArea struct {
	mOrm.Model
	CustomerID  uint     `gorm:"index;not null"`
	CountryCode string   `gorm:"size:3;index"`
	Zones       []string `gorm:"serializer:json"`
}
