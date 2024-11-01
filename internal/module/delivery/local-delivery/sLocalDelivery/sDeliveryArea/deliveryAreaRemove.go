package sDeliveryArea

import "shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"

func (s *sDeliveryArea) DeliveryAreaRemove(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}

	// 删除区域
	if err := s.orm.Model(&mLocalDelivery.LocalDeliveryArea{}).
		Where("shop_id = ? AND id = ?", s.shopId, ids).
		Delete(&mLocalDelivery.LocalDeliveryArea{}).Error; err != nil {
		return err
	}

	// 删除区域运费
	return s.orm.Model(&mLocalDelivery.LocalDeliveryFee{}).
		Where("shop_id = ? AND local_delivery_area_id IN ?", s.shopId, ids).
		Delete(&mLocalDelivery.LocalDeliveryFee{}).Error
}
