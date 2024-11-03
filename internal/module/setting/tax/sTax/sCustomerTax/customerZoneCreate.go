package sCustomerTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sCustomerTax) CustomerZoneCreate(zones []mTax.CustomerTaxZone) (err error) {
	if len(zones) == 0 {
		return err
	}
	zones = slice.Map(zones, func(index int, item mTax.CustomerTaxZone) mTax.CustomerTaxZone {
		item.ID = 0
		item.ShopId = s.shopId
		return item
	})
	return s.orm.Create(&zones).Error
}
