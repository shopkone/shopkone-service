package sShippingZone

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShippingZone) CodesRemove(zoneIds []uint) error {
	var codes []mShipping.ShippingZoneCode
	if err := s.orm.Model(&codes).Where("shipping_zone_id IN ?", zoneIds).
		Where("shop_id = ?", s.shopId).
		Delete(&codes).Error; err != nil {
		return err
	}
	return nil
}
