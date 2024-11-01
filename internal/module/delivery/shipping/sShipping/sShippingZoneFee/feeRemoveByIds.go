package sShippingZoneFee

import (
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

func (s *sShippingZoneFee) FeeRemoveByIds(feeIds []uint) (err error) {
	if len(feeIds) == 0 {
		return nil
	}
	return s.orm.Model(&mShipping.ShippingZoneFee{}).
		Where("id IN ? AND shop_id = ?", feeIds, s.shopId).
		Delete(&mShipping.ShippingZoneFee{}).Error
}
