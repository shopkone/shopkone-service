package sTax

import "shopkone-service/internal/module/setting/tax/mTax"

func (s *sTax) TaxZoneIsChange(oldZone, newZone mTax.TaxZone) bool {
	if oldZone.ZoneCode != newZone.ZoneCode {
		return true
	}
	if oldZone.Name != newZone.Name {
		return true
	}
	return oldZone.TaxRate != newZone.TaxRate
}
