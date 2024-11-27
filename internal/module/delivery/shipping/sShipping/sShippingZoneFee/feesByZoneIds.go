package sShippingZoneFee

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShippingZoneFee) FeesByZoneIds(ids []uint) (out []mShipping.ShippingZoneFee, err error) {
	if err = s.orm.Model(&out).
		Where("shipping_zone_id IN ?", ids).
		Omit("created_at", "deleted_at", "updated_at").
		Find(&out).Error; err != nil {
		return
	}
	return out, err
}
