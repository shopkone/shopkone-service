package sInventory

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/inventory/mInventory"
)

func (s *sInventory) MoveInventory(originLocationId, targetLocationId uint, email string) (err error) {
	// 找出来源仓库
	originInventories, err := s.GetInventoriesByLocationId(originLocationId)
	if err != nil {
		return err
	}
	// 转移的话，是转移源仓库库存不为0的
	originInventories = slice.Filter(originInventories, func(index int, item mInventory.Inventory) bool {
		return item.Quantity > 0
	})
	if len(originInventories) == 0 {
		return nil
	}
	originVariantIds := slice.Map(originInventories, func(index int, item mInventory.Inventory) uint {
		return item.VariantId
	})
	originVariantIds = slice.Unique(originVariantIds)
	originVariantIds = slice.Filter(originVariantIds, func(index int, item uint) bool {
		return item > 0
	})

	// 找出来目标库存
	targetInventories, err := s.ListByVariantsIds(originVariantIds, targetLocationId)
	if err != nil {
		return err
	}

	// 更新目标库存
	newTargetInventories := slice.Map(targetInventories, func(index int, item mInventory.Inventory) mInventory.Inventory {
		originInventory, ok := slice.FindBy(originInventories, func(i int, o mInventory.Inventory) bool {
			return o.VariantId == item.VariantId
		})
		if !ok {
			return item
		}
		item.Quantity += originInventory.Quantity
		return item
	})
	updateIn := iInventory.UpdateByDiffIn{
		News:        newTargetInventories,
		Olds:        targetInventories,
		HandleEmail: email,
		UpdateType:  mInventory.InventoryChangeTransfer,
	}
	if err = s.UpdateByDiff(updateIn); err != nil {
		return err
	}

	// 创建目标仓库中没有的产品
	createVariantIds := slice.Filter(originVariantIds, func(index int, item uint) bool {
		_, ok := slice.FindBy(newTargetInventories, func(i int, o mInventory.Inventory) bool {
			return o.VariantId == item
		})
		return !ok
	})

	// TODO: 将整个商品加入location中

	var createInventories []iInventory.CreateInventoryIn
	slice.ForEach(createVariantIds, func(index int, item uint) {
		find, ok := slice.FindBy(originInventories, func(index int, i mInventory.Inventory) bool {
			return i.VariantId == item
		})
		if !ok {
			return
		}
		temp := iInventory.CreateInventoryIn{
			VariantId:  item,
			LocationId: targetLocationId,
			Quantity:   find.Quantity,
		}
		createInventories = append(createInventories, temp)
	})
	if err = s.Create(createInventories, mInventory.InventoryChangeTransfer, email).Error; err != nil {
		return err
	}

	// 删除源仓库的库存
	if err = s.RemoveByVariantIds(originVariantIds, originLocationId); err != nil {
		return err
	}

	return nil
}
