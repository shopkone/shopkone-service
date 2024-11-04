package sTax

import (
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sTax) TaxZoneRemoveByTaxIds(taxIds []uint) (err error) {
	if len(taxIds) == 0 {
		return err
	}
	return s.orm.Where("tax_id IN (?) AND shop_id = ?", taxIds, s.shopId).
		Delete(&mTax.TaxZone{}).Error
}
