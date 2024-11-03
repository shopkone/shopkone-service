package sCustomerTax

import "shopkone-service/internal/module/setting/tax/mTax"

func (s *sCustomerTax) CustomerZoneRemove(customerIds []uint) (err error) {
	if len(customerIds) == 0 {
		return err
	}
	return s.orm.Where("shop_id = ? AND customer_tax_id IN ?", s.shopId, customerIds).
		Delete(&mTax.CustomerTaxZone{}).Error
}
