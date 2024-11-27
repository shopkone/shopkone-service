package sShipping

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShipping) ShippingGeneral() (out mShipping.Shipping, err error) {
	if err = s.orm.Model(&out).Where("type = ?", mShipping.GeneralExpressDelivery).
		Omit("created_at", "deleted_at", "updated_at").
		Find(&out).Error; err != nil {
		return out, err
	}
	return out, err
}
