package sInventory

import "shopkone-service/internal/module/product/inventory/mInventory"

func (s *sInventory) RemoveByVariantIds(variantIds []uint, locationId uint) error {
	query := s.orm.Where("variant_id in ? AND shop_id = ?", variantIds, s.shopId)
	if locationId != 0 {
		query = query.Where("location_id = ?", locationId)
	}
	return query.Unscoped().Delete(&mInventory.Inventory{}).Error
}
