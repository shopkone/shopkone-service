package sTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sTax) TaxZoneCreate(in []mTax.TaxZone) (err error) {
	if len(in) == 0 {
		return err
	}
	in = slice.Map(in, func(index int, item mTax.TaxZone) mTax.TaxZone {
		item.ShopId = s.shopId
		item.ID = 0
		return item
	})
	return s.orm.Create(&in).Error
}
