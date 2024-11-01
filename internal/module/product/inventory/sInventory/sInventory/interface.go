package sInventory

import (
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/utility/handle"
)

type sInventory struct {
	shopId uint
	orm    *gorm.DB
	data   []mInventory.Inventory
	Error  error
}

func NewInventory(orm *gorm.DB, shopId uint) *sInventory {
	return &sInventory{shopId: shopId, orm: orm}
}

type IInventory interface {
	// Create 创建库存
	Create(in []iInventory.CreateInventoryIn, inType mInventory.InventoryType, email string) (ids []uint, err error)
	// ListByVariantsIds 根据变体id获取库存列表
	ListByVariantsIds(variantIds []uint) ([]mInventory.Inventory, error)
	// RemoveByVariantIds 根据变体id删除库存
	RemoveByVariantIds(variantIds []uint) error
	// UpdateByDiff 更新库存
	UpdateByDiff(nin iInventory.UpdateByDiffIn) error
	// List 获取库存列表
	List(in vo.InventoryListReq) (out handle.PageRes[vo.InventoryListRes], err error)
	// CountByVariantIds 统计库存数量
	CountByVariantIds(variantIds []uint, locationId uint) (out []iInventory.CountByVariantIdsOut, err error)
	// CopyLocationInventory 复制仓库
	CopyLocationInventory(oldLocationId, newLocationId uint) (err error)
	// MoveInventory 转移库存
	MoveInventory(originLocationId, targetLocationId uint, email string) (err error)
	// ExistQuantityByLocationId 检查库存是否为0
	ExistQuantityByLocationId(locationId uint) (bool, error)
	// InventoryListUnByVariantIds 获取被删除的库存列表
	InventoryListUnByVariantIds(variantIds []uint) ([]mInventory.Inventory, error)
	// 强制删除库存
	ForceDeleteByVariantIds(variantIds []uint) error
}
