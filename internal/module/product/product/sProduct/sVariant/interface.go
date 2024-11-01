package sVariant

import (
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/product/iProduct"
	"shopkone-service/internal/module/product/product/mProduct"
)

type IVariant interface {
	// Create 创建变体
	Create(in iProduct.VariantCreateIn) (ids []uint, err error)
	// Update 更新变体
	Update(in iProduct.VariantUpdateIn) error
	// ListByProductId 根据商品ID获取变体列表
	ListByProductId(productId uint) ([]vo.BaseVariant, error)
	// GenInventoriesByVariants 根据变体列表生成库存列表
	GenInventoriesByVariants(variants []mProduct.Variant, in []vo.BaseVariant) (res []iInventory.CreateInventoryIn, err error)
	// IsChanged 判断是否变更
	IsChanged(newVariant, oldVariant mProduct.Variant) bool
	// UpdateInventory 更新库存
	UpdateInventory(variants []mProduct.Variant, newVariants []vo.BaseVariantWithId, trackInventory bool, email string) error
	// ListByProductIds 根据商品ID列表获取变体列表
	ListByProductIds(productIds []uint) ([]vo.BaseVariant, error)
	// ListByIds 根据变体ID获取变体
	ListByIds(ids []uint) (res []iProduct.VariantListByIdOut, err error)
}

type sVariant struct {
	orm    *gorm.DB
	shopId uint
}

func NewVariant(orm *gorm.DB, shopId uint) *sVariant {
	return &sVariant{orm, shopId}
}
