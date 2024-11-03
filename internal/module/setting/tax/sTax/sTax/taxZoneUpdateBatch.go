package sTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/utility/handle"
)

func (s *sTax) TaxZoneUpdateBatch(data, oldZones []mTax.TaxZone) (err error) {
	data = slice.Filter(data, func(index int, newZone mTax.TaxZone) bool {
		find, ok := slice.FindBy(oldZones, func(index int, oldZone mTax.TaxZone) bool {
			return oldZone.ID == newZone.ID
		})
		if !ok {
			return false
		}
		return s.TaxZoneIsChange(find, newZone)
	})
	if len(data) == 0 {
		return err
	}
	data = slice.Map(data, func(index int, item mTax.TaxZone) mTax.TaxZone {
		item.ShopId = s.shopId
		item.CanCreateId = true
		return item
	})
	batchIn := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"zone_code", "name", "tax_rate"},
	}
	return handle.BatchUpdateById(batchIn, &data)
}
