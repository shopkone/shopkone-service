package sCustomerTax

import (
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sCustomerTax) CustomerZoneRemoveIds(ids []uint) (err error) {
	if len(ids) == 0 {
		return err
	}
	return s.orm.Where("shop_id = ? AND id IN ?", s.shopId, ids).
		Delete(&mTax.CustomerTaxZone{}).Error
}
