package sInventoryChanger

import (
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/mInventory"
)

type InventoryChange interface {
	// AddHistory 添加库存历史记录
	AddHistory(in []mInventory.InventoryChange) error
	// List 获取库存历史
	List(id uint) ([]vo.InventoryHistoryRes, error)
}

type sInventoryChange struct {
	orm    *gorm.DB
	shopId uint
}

func NewInventoryChange(orm *gorm.DB, shopId uint) *sInventoryChange {
	return &sInventoryChange{orm: orm, shopId: shopId}
}
