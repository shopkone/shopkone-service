package sShippingZone

import (
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZoneFee"
)

func (s *sShippingZone) ZoneRemove(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}

	// 删除区域
	var zones []mShipping.ShippingZone
	if err := s.orm.Model(&zones).Where("shop_id = ? AND id IN ?", s.shopId, ids).Delete(&zones).Error; err != nil {
		return err
	}

	// 删除运费方案
	if err := sShippingZoneFee.NewShippingZoneFee(s.orm, s.shopId).FeeRemove(ids); err != nil {
		return err
	}

	// 删除codes
	return s.CodesRemove(ids)
}
