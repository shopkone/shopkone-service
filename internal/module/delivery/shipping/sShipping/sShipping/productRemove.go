package sShipping

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShipping) ProductRemove(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}
	return s.orm.Model(&mShipping.ShippingProduct{}).Where("shop_id = ?", s.shopId).
		Where("product_id IN ?", ids).Delete(&mShipping.ShippingProduct{}).Error
}
