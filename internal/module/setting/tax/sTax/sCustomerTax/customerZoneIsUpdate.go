package sCustomerTax

import "shopkone-service/internal/module/setting/tax/mTax"

func (s *sCustomerTax) CustomerZoneIsUpdate(oldZone, newZone mTax.CustomerTaxZone) bool {
	if oldZone.AreaCode != newZone.AreaCode {
		return true
	}
	if oldZone.TaxRate != newZone.TaxRate {
		return true
	}
	if oldZone.Name != newZone.Name {
		return true
	}
	return false
}
