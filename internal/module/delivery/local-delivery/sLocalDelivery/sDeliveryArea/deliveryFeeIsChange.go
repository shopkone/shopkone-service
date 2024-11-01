package sDeliveryArea

import "shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"

func (s *sDeliveryArea) DeliveryFeeIsChange(newItem, oldItem mLocalDelivery.LocalDeliveryFee) bool {
	if newItem.Condition != oldItem.Condition || newItem.Fee != oldItem.Fee {
		return true
	}
	return false
}
