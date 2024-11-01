package mTax

import "shopkone-service/internal/module/base/orm/mOrm"

type TaxStatus uint8

const (
	Active   TaxStatus = 1
	Inactive TaxStatus = 2
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
