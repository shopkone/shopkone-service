package mProduct

import (
	"shopkone-service/internal/consts"
	"shopkone-service/internal/module/base/orm/mOrm"
)

type VariantNameHandler struct {
	mOrm.Model
	ProductId uint
	Value     string
	Label     string
	VariantId uint
}

// 商品变体名称
type VariantName struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// 商品变体
type Variant struct {
	mOrm.Model
	Price            uint32            `gorm:"not null"`        // 价格
	CostPerItem      *uint32           `gorm:"default:null"`    // 每件成本
	CompareAtPrice   *uint32           `gorm:"default:null"`    // 原价
	WeightUnit       consts.WeightUnit `gorm:"default:g"`       // 重量单位
	Weight           *float32          `gorm:"default:null"`    // 重量
	Sku              string            `gorm:"size:250"`        // SKU
	Barcode          string            `gorm:"size:250"`        // 条形码
	ProductId        uint              `gorm:"index;not null"`  // 商品ID
	ImageId          uint              `gorm:"index;not null"`  // 图片ID
	Name             []VariantName     `gorm:"serializer:json"` // 名称
	TaxRequired      bool              `gorm:"default:0"`       // 税是否必须
	ShippingRequired bool              `gorm:"default:0"`       // 快递是否必须
}
