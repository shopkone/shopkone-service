package sShippingZoneFee

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShippingZoneFee) FeeRemove(zoneIds []uint) (err error) {
	if err = s.orm.Model(mShipping.ShippingZoneFee{}).
		Where("shipping_zone_id IN ? AND shop_id = ?", zoneIds, s.shopId).
		Delete(&mShipping.ShippingZoneFee{}).Error; err != nil {
		return err
	}
	return nil
}
