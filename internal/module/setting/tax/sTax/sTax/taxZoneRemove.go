package sTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sTax) TaxZoneRemove(remove []mTax.TaxZone) (err error) {
	removeIds := slice.Map(remove, func(index int, item mTax.TaxZone) uint { return item.ID })
	if len(removeIds) == 0 {
		return err
	}
	return s.orm.Where("id IN (?) AND shop_id = ?", removeIds, s.shopId).
		Delete(&mTax.TaxZone{}).Error
}
