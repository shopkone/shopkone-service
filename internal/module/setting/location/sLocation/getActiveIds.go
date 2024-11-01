package sLocation

import "shopkone-service/internal/module/setting/location/mLocation"

func (s *sLocation) GetActiveIds() ([]uint, error) {
	var locations []mLocation.Location
	if err := s.orm.Where("shop_id = ? AND active = ?", s.shopId, true).
		Select("id").Order("order_num ASC").Find(&locations).Error; err != nil {
		return nil, err
	}

	ids := make([]uint, 0, len(locations))
	for _, location := range locations {
		ids = append(ids, location.ID)
	}
	return ids, nil
}
