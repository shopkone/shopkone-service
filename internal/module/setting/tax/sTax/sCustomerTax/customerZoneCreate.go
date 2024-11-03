package sCustomerTax

import (
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/utility/code"
)

func (s *sCustomerTax) CustomerZoneCreate(zones []mTax.CustomerTaxZone) (err error) {
	if len(zones) == 0 {
		return code.TaxCustomerZonesMust
	}
	return s.orm.Create(&zones).Error
}
