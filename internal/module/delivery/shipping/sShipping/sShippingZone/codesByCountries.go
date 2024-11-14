package sShippingZone

import "shopkone-service/internal/module/delivery/shipping/mShipping"

func (s *sShippingZone) CodesByCountries(countryCodes []string) (res []mShipping.ShippingZoneCode, err error) {
	query := s.orm.Model(&res).Where("country_code IN ?", countryCodes)
	query = query.Omit("created_at", "updated_at", "deleted_at", "shop_id")
	return res, query.Find(&res).Error
}
