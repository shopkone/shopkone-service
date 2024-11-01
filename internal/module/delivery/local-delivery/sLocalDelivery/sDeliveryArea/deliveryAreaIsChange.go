package sDeliveryArea

import "shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"

func (s *sDeliveryArea) DeliveryAreaIsChange(o, n mLocalDelivery.LocalDeliveryArea) bool {
	if o.Note != n.Note {
		return true
	}
	if o.Name != n.Name {
		return true
	}
	if o.PostalCode != n.PostalCode {
		return true
	}
	return false
}
