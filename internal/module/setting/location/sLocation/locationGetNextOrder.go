package sLocation

import "shopkone-service/internal/module/setting/location/mLocation"

func (s *sLocation) LocationGetNextOrder() (count int64, err error) {
	query := s.orm.Model(&mLocation.Location{}).Where("shop_id = ?", s.shopId)
	query = query.Unscoped()
	if err = query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count + 1, err
}
