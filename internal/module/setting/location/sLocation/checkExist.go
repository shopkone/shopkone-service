package sLocation

import (
	"shopkone-service/internal/module/setting/location/mLocation"
	"shopkone-service/utility/code"
)

func (s *sLocation) CheckExist(id uint) error {
	var count int64
	if err := s.orm.Model(&mLocation.Location{}).Where("id = ? AND active = ?", id, true).
		Where("shop_id = ?", s.shopId).Count(&count).Error; err != nil {
		return err
	}
	if count <= 0 {
		return code.ErrLocationNotFound
	}
	return nil
}
