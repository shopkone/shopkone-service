package mMarket

import "shopkone-service/internal/module/base/orm/mOrm"

type PriceAdjustmentType uint

const (
	PriceAdjustmentTypeReduce PriceAdjustmentType = iota + 1 // 减少
	PriceAdjustmentTypeAdd                                   // 增加
)

// 定价调整
type MarketPrice struct {
	mOrm.Model
	MarketID      uint                `gorm:"index"`
	CurrencyCode  string              `gorm:"index;not null;default:USD'"` // 货币
	AdjustPercent float64             `json:"percent"`
	AdjustType    PriceAdjustmentType `gorm:"default:1"`
}

// 商品调整
type MarketProduct struct {
	mOrm.Model
	MarketID  uint     `gorm:"index;uniqueIndex:id_product"`
	ProductID uint     `gorm:"index;uniqueIndex:id_product"` // 商品层面进行调整
	Fixed     *float64 `gorm:"index"`                        // 固定金额
	Exclude   bool     `gorm:"default:false"`                // 是否排除
}
