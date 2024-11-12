package mMarket

import "shopkone-service/internal/module/base/orm/mOrm"

// 商品调整
type MarketProduct struct {
	mOrm.Model
	MarketID  uint     `gorm:"index;uniqueIndex:id_product"`
	ProductID uint     `gorm:"index;uniqueIndex:id_product"` // 商品层面进行调整
	Fixed     *float64 `gorm:"index"`                        // 固定金额
	Exclude   bool     `gorm:"default:false"`                // 是否排除
}
