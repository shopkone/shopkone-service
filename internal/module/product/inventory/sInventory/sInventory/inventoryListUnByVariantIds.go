package sInventory

import "shopkone-service/internal/module/product/inventory/mInventory"

func (s *sInventory) InventoryListUnByVariantIds(variantIds []uint, locationIds []uint) (out []mInventory.Inventory, err error) {
	query := s.orm.Model(out).Unscoped()
	query = query.Where("shop_id = ? AND variant_id IN ?", s.shopId, variantIds)
	query = query.Where("location_id IN ?", locationIds)
	return out, query.Find(&out).Error
}
