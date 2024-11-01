package mInventory

import "shopkone-service/internal/module/base/orm/mOrm"

// 库存
type Inventory struct {
	mOrm.Model
	Quantity   uint `gorm:"not null"`                                        // 库存数量
	LocationId uint `gorm:"index;not null;uniqueIndex:idx_location_variant"` // 地点Id
	VariantId  uint `gorm:"index;not null;uniqueIndex:idx_location_variant"` // 商品变体Id
}
