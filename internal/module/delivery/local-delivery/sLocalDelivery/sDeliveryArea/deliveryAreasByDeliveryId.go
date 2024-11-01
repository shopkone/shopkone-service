package sDeliveryArea

import "shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"

func (s *sDeliveryArea) DeliveryAreasByDeliveryId(deliveryId uint) ([]mLocalDelivery.LocalDeliveryArea, error) {
	var areas []mLocalDelivery.LocalDeliveryArea
	return areas, s.orm.Model(&mLocalDelivery.LocalDeliveryArea{}).
		Where("shop_id = ? AND local_delivery_id = ?", s.shopId, deliveryId).
		Select("id").Find(&areas).Error
}
