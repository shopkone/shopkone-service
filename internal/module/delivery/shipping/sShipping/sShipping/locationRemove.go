package sShipping

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShipping) LocationRemove(shippingId uint, locationIds []uint) error {
	if len(locationIds) == 0 {
		return nil
	}
	return s.orm.Model(&mShipping.ShippingLocation{}).
		Where("shipping_id = ? AND location_id IN ? AND shop_id = ?", shippingId, locationIds, s.shopId).
		Delete(&mShipping.ShippingLocation{}).Error
}
