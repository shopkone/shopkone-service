package sShippingZoneFee

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShippingZoneFee) ConditionsByFeeIds(feeIds []uint) (out []mShipping.ShippingZonFeeCondition, err error) {
	if err = s.orm.Model(&out).
		Where("shipping_zone_fee_id IN ? AND shop_id = ?", feeIds, s.shopId).
		Omit("created_at", "deleted_at", "updated_at").
		Find(&out).Error; err != nil {
		return
	}
	return out, err
}
