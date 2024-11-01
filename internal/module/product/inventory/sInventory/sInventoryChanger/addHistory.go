package sInventoryChanger

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/inventory/mInventory"
)

func (s *sInventoryChange) AddHistory(in []mInventory.InventoryChange) error {
	if len(in) == 0 {
		return nil
	}
	in = slice.Map(in, func(index int, item mInventory.InventoryChange) mInventory.InventoryChange {
		item.ShopId = s.shopId
		item.ID = 0
		return item
	})
	return s.orm.Create(&in).Error
}
