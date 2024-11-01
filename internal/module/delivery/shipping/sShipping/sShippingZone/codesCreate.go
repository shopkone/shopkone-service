package sShippingZone

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShippingZone) CodesCreate(codes []mShipping.ShippingZoneCode) error {
	if err := s.orm.Create(&codes).Error; err != nil {
		return err
	}
	return s.CodesUpdateTaxs(codes)
}
