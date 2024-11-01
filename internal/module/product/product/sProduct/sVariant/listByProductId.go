package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/product/product/mProduct"
	sTransfer2 "shopkone-service/internal/module/product/product/sProduct/sTransfer"
)

func (s *sVariant) ListByProductId(productId uint) ([]vo.BaseVariant, error) {
	// 获取变体列表
	var list []mProduct.Variant
	err := s.orm.Where("product_id = ?", productId).Find(&list).Error
	if err != nil {
		return nil, err
	}
	// 获取库存列表
	variantIds := slice.Map(list, func(index int, item mProduct.Variant) uint {
		return item.ID
	})
	inventoryList, err := sInventory.NewInventory(s.orm, s.shopId).ListByVariantsIds(variantIds, 0)
	if err != nil {
		return nil, err
	}
	// 组装数据
	sTransfer := sTransfer2.NewProductTransfer(s.shopId)
	result := slice.Map(list, func(index int, item mProduct.Variant) vo.BaseVariant {
		// 组装库存
		variantInventories := slice.Filter(inventoryList, func(index int, i mInventory.Inventory) bool {
			return i.VariantId == item.ID
		})
		inventories := slice.Map(variantInventories, func(index int, i mInventory.Inventory) vo.VariantInventory {
			return vo.VariantInventory{
				Id:         i.ID,
				LocationId: i.LocationId,
				Quantity:   i.Quantity,
			}
		})
		// 组装返回
		temp := sTransfer.ModelToVariant(item)
		temp.Inventories = inventories
		return temp
	})
	return result, err
}
