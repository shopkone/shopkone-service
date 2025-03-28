package mTax

import "shopkone-service/internal/module/base/orm/mOrm"

type TaxStatus uint8

const (
	TaxStatusActive   TaxStatus = 1
	TaxStatusInactive TaxStatus = 2
)

type CustomerTaxType uint8

const (
	CustomerTaxTypeCollection CustomerTaxType = 1 // 自定义-系列
	CustomerTaxTypeDelivery   CustomerTaxType = 2 // 自定义-配送
)

type Tax struct {
	mOrm.Model
	CountryCode string    `gorm:"index;not null"`
	TaxRate     float64   `gorm:"default:0"`     // 税率
	HasNote     bool      `gorm:"default:false"` // 是否有备注
	Note        string    `gorm:"size:500"`      // 备注
	Status      TaxStatus `gorm:"default:2"`     // 状态
}

type TaxZone struct {
	mOrm.Model
	ZoneCode string  `gorm:"index;not null"`
	Name     string  `gorm:"size:200"`
	TaxRate  float64 `gorm:"default:0"`
	TaxId    uint    `gorm:"index;not null"`
}

type CustomerTax struct {
	mOrm.Model
	Type         CustomerTaxType
	TaxID        uint `gorm:"index;not null"`
	CollectionID uint `gorm:"index"`
}

type CustomerTaxZone struct {
	mOrm.Model
	AreaCode      string  `gorm:"index;not null"` // 国家代码 + 省份代码
	Name          string  `gorm:"size:200"`
	TaxRate       float64 `gorm:"default:0"`
	CustomerTaxID uint    `gorm:"index;not null"`
}
