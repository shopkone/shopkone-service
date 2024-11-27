package sCustomerTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/utility/handle"
)

func (s *sCustomerTax) TaxZoneUpdateBatch(list, oldList []mTax.CustomerTaxZone) (err error) {
	list = slice.Filter(list, func(index int, item mTax.CustomerTaxZone) bool {
		old, ok := slice.FindBy(oldList, func(index int, old mTax.CustomerTaxZone) bool {
			return old.ID == item.ID
		})
		if !ok {
			return false
		}
		return s.CustomerZoneIsUpdate(old, item)
	})
	if len(list) == 0 {
		return err
	}
	list = slice.Map(list, func(index int, item mTax.CustomerTaxZone) mTax.CustomerTaxZone {
		item.ShopId = s.shopId
		item.CanCreateId = true
		return item
	})
	batchIn := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"area_code", "tax_rate", "name"},
	}
	return handle.BatchUpdateById(batchIn, &list)
}
