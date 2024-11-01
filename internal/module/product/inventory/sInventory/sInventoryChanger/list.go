package sInventoryChanger

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/utility/handle"
)

func (s *sInventoryChange) List(id uint) ([]vo.InventoryHistoryRes, error) {
	var list []mInventory.InventoryChange
	err := s.orm.Model(&mInventory.InventoryChange{}).Where("inventory_id = ?", id).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return slice.Map(list, func(index int, item mInventory.InventoryChange) vo.InventoryHistoryRes {
		i := vo.InventoryHistoryRes{}
		i.Date = handle.ToUnix(item.CreatedAt)
		i.Id = item.ID
		i.Activity = item.Type
		i.DiffQuantity = item.DiffQuantity
		return i
	}), nil
}
