package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/setting/location/sLocation"
	"shopkone-service/utility/code"
)

type GenIn struct {
	Variants           []mProduct.Variant
	Base               []vo.BaseVariant
	EnabledLocationIds []uint
}

func (s *sVariant) GenInventoriesByVariants(in GenIn) (res []iInventory.CreateInventoryIn, err error) {
	if len(in.Variants) == 0 || len(in.Base) == 0 {
		return res, nil
	}
	if len(in.EnabledLocationIds) == 0 {
		return res, code.IdMissing
	}
	variants := in.Variants
	base := in.Base
	// 获取现存可用的地点
	locationIds, err := sLocation.NewLocation(s.orm, s.shopId).GetActiveIds()
	locationIds = slice.Filter(locationIds, func(index int, item uint) bool {
		_, ok := slice.FindBy(in.EnabledLocationIds, func(index int, id uint) bool {
			return id == item
		})
		return ok
	})
	if err != nil {
		return res, err
	}
	if len(locationIds) == 0 {
		return res, err
	}
	// 生成
	slice.ForEach(variants, func(index int, item mProduct.Variant) {
		find, ok := slice.FindBy(base, func(index int, inItem vo.BaseVariant) bool {
			name1 := slice.Map(item.Name, func(index int, ni mProduct.VariantName) string {
				return ni.Value + ni.Label
			})
			name2 := slice.Map(inItem.Name, func(index int, ni mProduct.VariantName) string {
				return ni.Value + ni.Label
			})
			return slice.Equal(name1, name2)
		})
		if !ok {
			return
		}
		slice.ForEach(locationIds, func(index int, locationId uint) {
			temp := iInventory.CreateInventoryIn{}
			temp.VariantId = item.ID
			temp.LocationId = locationId
			vi, ok := slice.FindBy(find.Inventories, func(index int, i vo.VariantInventory) bool {
				return i.LocationId == locationId
			})
			if ok {
				temp.Id = vi.Id
				temp.Quantity = vi.Quantity
			}
			res = append(res, temp)
		})
	})
	return res, err
}
