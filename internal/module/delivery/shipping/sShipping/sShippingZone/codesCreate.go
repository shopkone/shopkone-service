package sShippingZone

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShippingZone) CodesCreate(codes []mShipping.ShippingZoneCode) error {
	return s.orm.Create(&codes).Error
}
