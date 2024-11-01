package sInventory

import "shopkone-service/internal/module/product/inventory/mInventory"

func (s *sInventory) GetInventoriesByLocationId(locationId uint) ([]mInventory.Inventory, error) {
	var inventories []mInventory.Inventory
	if err := s.orm.Where("shop_id =? AND location_id = ?", s.shopId, locationId).
		Omit("created_at", "updated_at", "deleted_at").Find(&inventories).Error; err != nil {
		return nil, err
	}
	return inventories, nil
}
