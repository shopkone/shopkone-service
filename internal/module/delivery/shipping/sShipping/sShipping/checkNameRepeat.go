package sShipping

import (
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/utility/code"
)

func (s *sShipping) CheckNameRepeat(name string, id uint) error {
	var oldCount int64
	if err := s.orm.Model(&mShipping.Shipping{}).
		Where("name = ? AND shop_id = ?", name, s.shopId).
		Where("id != ?", id).
		Count(&oldCount).Error; err != nil {
		return err
	}
	if oldCount > 0 {
		return code.ShippingNameExist
	}
	return nil
}
