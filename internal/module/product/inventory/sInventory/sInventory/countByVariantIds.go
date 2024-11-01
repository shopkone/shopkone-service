package sInventory

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/inventory/mInventory"
)

func (s *sInventory) CountByVariantIds(variantIds []uint, locationId uint) (out []iInventory.CountByVariantIdsOut, err error) {
	var list []mInventory.Inventory
	query := s.orm.Model(&list).Where("shop_id = ?", s.shopId)
	query = query.Where("variant_id IN (?)", variantIds)
	if locationId != 0 {
		query = query.Where("location_id = ?", locationId)
	}
	query = query.Select("quantity", "variant_id")
	if err = query.Find(&list).Error; err != nil {
		return out, err
	}
	out = slice.Map(variantIds, func(index int, item uint) iInventory.CountByVariantIdsOut {
		inventories := slice.Filter(list, func(index int, i mInventory.Inventory) bool {
			return i.VariantId == item
		})
		var quantity uint
		slice.ForEach(inventories, func(index int, ii mInventory.Inventory) {
			quantity += ii.Quantity
		})
		return iInventory.CountByVariantIdsOut{
			Quantity:  quantity,
			VariantId: item,
		}
	})
	return out, err
}
