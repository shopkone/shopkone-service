package sShippingZoneFee

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShippingZoneFee) FeeIsChange(newFee, oldFee mShipping.ShippingZoneFee) bool {
	if newFee.Name != oldFee.Name {
		return true
	}
	if newFee.WeightUnit != oldFee.WeightUnit {
		return true
	}
	if newFee.Type != oldFee.Type {
		return true
	}
	if newFee.Rule != oldFee.Rule {
		return true
	}
	if newFee.CurrencyCode != oldFee.CurrencyCode {
		return true
	}
	if newFee.Remark != oldFee.Remark {
		return true
	}
	if newFee.Cod != oldFee.Cod {
		return true
	}
	return false
}
