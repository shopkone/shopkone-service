package sLocation

import "shopkone-service/internal/module/setting/location/mLocation"

func (s *sLocation) LocationSetDefault(locationId uint) error {
	// 将旧的默认位置设置为非默认
	if err := s.orm.Model(&mLocation.Location{}).
		Where("shop_id = ?", s.shopId).
		Where("is_default = ?", true).
		Update("is_default", false).Error; err != nil {
		return err
	}

	query := s.orm.Model(&mLocation.Location{})
	query = query.Where("shop_id = ?", s.shopId)
	query = query.Where("active = ?", true)
	query = query.Where("id = ?", locationId)
	return query.Update("is_default", true).Error
}
