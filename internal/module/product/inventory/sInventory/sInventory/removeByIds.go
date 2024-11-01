package sInventory

import "shopkone-service/internal/module/product/inventory/mInventory"

func (s *sInventory) RemoveByIds(ids []uint) error {
	return s.orm.Where("id in? AND shop_id =?", ids, s.shopId).
		Unscoped().Delete(&mInventory.Inventory{}).Error
}
