package sInStorePickup

import "shopkone-service/internal/module/delivery/in-store-pick-up/mInStorePickup"

func (s *sInStorePickup) RemoveByLocationId(locationId uint) error {
	// 获取到店自提
	var info mInStorePickup.InStorePickup
	if err := s.orm.Where("location_id = ? AND shop_id = ?", locationId, s.shopId).
		Select("id").First(&info).Error; err != nil {
		return err
	}

	// 删除到店自提
	if err := s.orm.Where("id = ? AND shop_id = ?", info.ID, s.shopId).
		Delete(&mInStorePickup.InStorePickup{}).Error; err != nil {
		return err
	}

	// 删除营业时间
	return s.orm.Model(&mInStorePickup.InStorePickupBusinessHours{}).
		Where("in_store_pickup_id = ? AND shop_id = ?", info.ID, s.shopId).
		Delete(&mInStorePickup.InStorePickupBusinessHours{}).Error
}
