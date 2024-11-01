package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/product/product/mProduct"
)

// 仅在更新变体时使用该函数
func (s *sVariant) UpdateInventoryWithUpdateVariants(variants []mProduct.Variant, newVariants []vo.BaseVariantWithId, trackInventory bool, email string) error {
	if !trackInventory {
		return nil
	}
	// 获取旧的库存
	updateIds := slice.Map(variants, func(index int, item mProduct.Variant) uint {
		return item.ID
	})
	si := sInventory.NewInventory(s.orm, s.shopId)
	oldInventories, err := si.ListByVariantsIds(updateIds, 0)
	if err != nil {
		return err
	}
	// 新的库存
	var newInventories []mInventory.Inventory
	slice.ForEach(variants, func(index int, i mProduct.Variant) {
		find, ok := slice.FindBy(newVariants, func(index int, item vo.BaseVariantWithId) bool {
			return item.Id == i.ID
		})
		if !ok {
			return
		}
		if find.Inventories != nil && len(find.Inventories) > 0 {
			slice.ForEach(find.Inventories, func(index int, inventory vo.VariantInventory) {
				temp := mInventory.Inventory{
					Quantity:   inventory.Quantity,
					LocationId: inventory.LocationId,
					VariantId:  i.ID,
				}
				temp.ID = inventory.Id
				temp.ShopId = s.shopId
				newInventories = append(newInventories, temp)
			})
		}
	})
	// 处理新的库存
	newInventories = slice.Map(newInventories, func(index int, item mInventory.Inventory) mInventory.Inventory {
		newInventory, ok := slice.FindBy(oldInventories, func(index int, old mInventory.Inventory) bool {
			return item.VariantId == old.VariantId && item.LocationId == old.LocationId
		})
		if !ok {
			return item
		}
		item.ID = newInventory.ID
		return item
	})
	// 更新
	updateInventoryIn := iInventory.UpdateByDiffIn{
		News:        newInventories,
		Olds:        oldInventories,
		HandleEmail: email,
		UpdateType:  mInventory.InventoryChangeProduct,
	}
	// 更新库存
	if err = si.UpdateByDiff(updateInventoryIn); err != nil {
		return err
	}
	return err
}
