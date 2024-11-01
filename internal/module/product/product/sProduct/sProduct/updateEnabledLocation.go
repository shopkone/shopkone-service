package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/setting/location/sLocation"
	"shopkone-service/utility/code"
)

// 更新使用地点
func (s *sProduct) UpdateEnabledLocationIds(productId uint, locationIds []uint, email string) error {
	activeLocationId, err := sLocation.NewLocation(s.orm, s.shopId).GetActiveIds()
	locationIds = slice.Filter(locationIds, func(index int, item uint) bool {
		return slice.Contain(activeLocationId, item)
	})

	// 获取该商品下的所有变体
	var variants []mProduct.Variant
	if err := s.orm.Model(&variants).Where("product_id = ? AND shop_id = ?", productId, s.shopId).
		Select("id").Find(&variants).Error; err != nil {
		return err
	}
	variantIds := slice.Map(variants, func(index int, item mProduct.Variant) uint {
		return item.ID
	})
	if len(variantIds) == 0 {
		return code.IdMissing
	}

	inventoryService := sInventory.NewInventory(s.orm, s.shopId)

	// 获取变体的所有库存
	inventories, err := inventoryService.ListByVariantsIds(variantIds, 0)
	if err != nil {
		return err
	}
	oldLocationIds := slice.Map(inventories, func(index int, item mInventory.Inventory) uint {
		return item.LocationId
	})
	oldLocationIds = slice.Unique(oldLocationIds)

	// 找出要删除的库存
	removeInventories := slice.Filter(inventories, func(index int, item mInventory.Inventory) bool {
		return !slice.Contain(locationIds, item.LocationId)
	})
	removeInventoryIds := slice.Map(removeInventories, func(index int, item mInventory.Inventory) uint {
		return item.ID
	})
	// 删除
	if err = inventoryService.RemoveByIds(removeInventoryIds); err != nil {
		return err
	}

	// 找出要添加的库存
	var createInventories []iInventory.CreateInventoryIn
	slice.ForEach(locationIds, func(index int, locationId uint) {
		slice.ForEach(variantIds, func(index int, variantId uint) {
			_, ok := slice.FindBy(inventories, func(index int, i mInventory.Inventory) bool {
				return i.VariantId == variantId && i.LocationId == locationId
			})
			if ok {
				return
			}
			i := iInventory.CreateInventoryIn{}
			i.VariantId = variantId
			i.LocationId = locationId
			createInventories = append(createInventories, i)
		})
	})
	if len(createInventories) == 0 {
		return err
	}

	if inventoryService.Create(createInventories, mInventory.InventoryChangeProduct, email); err != nil {
		return err
	}

	return err
}
