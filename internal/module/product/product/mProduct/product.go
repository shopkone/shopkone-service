package mProduct

import (
	"shopkone-service/internal/module/base/orm/mOrm"
	"time"
)

/*InventoryPolicy 商品缺货库存策略*/
type InventoryPolicy uint8

const (
	InventoryPolicyContinue InventoryPolicy = iota + 1 // 仍可下单
	InventoryPolicyDeny                                // 禁止下单
	InventoryPolicyDownward                            // 下架
)

/*VariantType 商品变体类型*/
type VariantType uint8

const (
	VariantTypeSingle VariantType = iota + 1 // 单一变体
	VariantTypeMulti                         // 多个变体
)

// VariantStatus 商品状态
type VariantStatus uint8

const (
	VariantStatusDraft     VariantStatus = 1 // 草稿
	VariantStatusPublished VariantStatus = 2 // 已发布
)

type Product struct {
	mOrm.Model
	Title             string          `gorm:"size:200"`
	Description       string          `gorm:"type:text"`
	Status            VariantStatus   `gorm:"default:1"`
	PublishedAt       *time.Time      `gorm:"default:null"`
	ScheduledAt       *time.Time      `gorm:"default:null"`
	InventoryTracking bool            `gorm:"default:1"`
	Spu               string          `gorm:"size:200"`
	Vendor            string          `gorm:"size:200"`
	Tags              []string        `gorm:"serializer:json"`
	SeoId             uint            `gorm:"index;not null"`
	VariantType       VariantType     `gorm:"default:1"`
	InventoryPolicy   InventoryPolicy `gorm:"default:1"`
	Category          uint
}

type ProductFiles struct {
	mOrm.Model
	ProductId uint `gorm:"index;not null"`
	FileId    uint `gorm:"index;not null"`
	Position  uint `gorm:"index;not null"`
}

type ProductOption struct {
	mOrm.Model
	ProductId uint     `gorm:"index;not null"`
	Label     string   `gorm:"index;not null"`
	Values    []string `gorm:"serializer:json"`
	ImageId   uint     `gorm:"index"`
}
