package sShippingZone

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/setting/tax/sTax/sTax"
)

func (s *sShippingZone) CodesUpdateTaxs(codes []mShipping.ShippingZoneCode) error {
	countryCodes := slice.Map(codes, func(index int, item mShipping.ShippingZoneCode) string {
		return item.CountryCode
	})
	countryCodes = slice.Unique(countryCodes)
	return sTax.NewTax(s.orm, s.shopId).TaxCreate(countryCodes)
}
