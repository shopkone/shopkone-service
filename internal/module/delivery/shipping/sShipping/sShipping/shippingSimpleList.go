package sShipping

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShipping) SimpleList(ids []uint) (res []mShipping.Shipping, err error) {
	query := s.orm.Model(&res).Where("id in (?)", ids)
	return res, query.Select("id", "name").Find(&res).Error
}
