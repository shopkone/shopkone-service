package sInventory

import (
	"github.com/duke-git/lancet/v2/slice"
	"gorm.io/gorm/clause"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventoryChanger"
)

func (s *sInventory) UpdateByDiff(nin iInventory.UpdateByDiffIn) error {
	olds := nin.Olds
	news := nin.News
	handleEmail := nin.HandleEmail
	updateType := nin.UpdateType
	// 获取需要更新的库存,只更新不相等的库存
	updatePart := slice.Filter(olds, func(index int, item mInventory.Inventory) bool {
		_, ok := slice.FindBy(news, func(index int, i mInventory.Inventory) bool {
			return item.ID == i.ID && item.Quantity != i.Quantity
		})
		return ok
	})
	// 批量更新前的处理
	updatePart = slice.Map(updatePart, func(index int, item mInventory.Inventory) mInventory.Inventory {
		find, ok := slice.FindBy(news, func(index int, i mInventory.Inventory) bool {
			return item.ID == i.ID
		})
		if !ok {
			return item
		}
		item.Quantity = find.Quantity
		item.CanCreateId = true
		return item
	})
	// 更新多条库存
	if len(updatePart) > 0 {
		if err := s.orm.Model(&updatePart).
			Where("shop_id = ?", s.shopId).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"quantity"}),
		}).Create(&updatePart).Error; err != nil {
			return err
		}
	}
	// 添加库存历史
	histories := slice.Map(updatePart, func(index int, item mInventory.Inventory) mInventory.InventoryChange {
		temp := mInventory.InventoryChange{}
		temp.HandleEmail = handleEmail
		temp.ShopId = s.shopId
		temp.InventoryId = item.ID
		temp.Type = updateType
		old, ok := slice.FindBy(olds, func(index int, i mInventory.Inventory) bool {
			return item.ID == i.ID
		})
		if !ok {
			return temp
		}
		temp.DiffQuantity = int(item.Quantity) - int(old.Quantity)
		if temp.DiffQuantity == 0 {
			return temp
		}
		if temp.DiffQuantity > 0 {
			temp.Style = mInventory.InventoryStyleIn
		}
		if temp.DiffQuantity < 0 {
			temp.Style = mInventory.InventoryStyleOut
		}
		return temp
	})
	histories = slice.Filter(histories, func(index int, item mInventory.InventoryChange) bool {
		return item.DiffQuantity != 0
	})
	if len(histories) > 0 {
		if err := sInventoryChanger.NewInventoryChange(s.orm, s.shopId).AddHistory(histories); err != nil {
			return err
		}
	}
	return nil
}
