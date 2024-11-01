package sInventory

import "shopkone-service/internal/module/product/inventory/mInventory"

func (s *sInventory) ListByVariantsIds(variantIds []uint, locationId uint) ([]mInventory.Inventory, error) {
	var res []mInventory.Inventory
	query := s.orm.Model(&mInventory.Inventory{}).Where("variant_id in ? AND shop_id = ?", variantIds, s.shopId)
	if locationId != 0 {
		query = query.Where("location_id = ?", locationId)
	}
	return res, query.Find(&res).Error
}
