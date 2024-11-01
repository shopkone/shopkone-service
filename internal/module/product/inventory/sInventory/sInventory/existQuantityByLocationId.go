package sInventory

import "shopkone-service/internal/module/product/inventory/mInventory"

func (s *sInventory) ExistQuantityByLocationId(locationId uint) (bool, error) {
	var count int64
	query := s.orm.Model(mInventory.Inventory{})
	query = query.Where("shop_id = ? AND location_id = ? AND quantity > 0", s.shopId, locationId)
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
