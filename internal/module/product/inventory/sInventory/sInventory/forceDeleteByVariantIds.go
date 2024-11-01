package sInventory

import "shopkone-service/internal/module/product/inventory/mInventory"

func (s *sInventory) ForceDeleteByVariantIds(variantIds []uint) error {
	query := s.orm.Model(mInventory.Inventory{})
	query = query.Where("shop_id = ? AND variant_id IN ?", s.shopId, variantIds)
	return query.Unscoped().Delete(&mInventory.Inventory{}).Error
}
