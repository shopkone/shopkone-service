package sShippingZoneFee

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShippingZoneFee) ConditionIsChange(oldCondition, newCondition mShipping.ShippingZonFeeCondition) (isChange bool) {
	if oldCondition.Fixed != newCondition.Fixed {
		return true
	}
	if oldCondition.First != newCondition.First {
		return true
	}
	if oldCondition.FirstFee != newCondition.FirstFee {
		return true
	}
	if oldCondition.Next != newCondition.Next {
		return true
	}
	if oldCondition.NextFee != newCondition.NextFee {
		return true
	}
	if oldCondition.Max != newCondition.Max {
		return true
	}
	if oldCondition.Min != newCondition.Min {
		return true
	}
	return false
}
