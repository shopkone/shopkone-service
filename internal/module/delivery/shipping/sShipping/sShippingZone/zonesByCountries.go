package sShippingZone

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

func (s *sShippingZone) ZonesByCountries(countryCodes []string) (res []vo.BaseShippingZone, err error) {
	// 获取国家地区
	zoneCodes, err := s.CodesByCountries(countryCodes)
	if err != nil {
		return nil, err
	}
	if len(zoneCodes) == 0 {
		return res, nil
	}
	// 晒出来区域id，获取区域数据
	var zoneIds []uint
	slice.ForEach(zoneCodes, func(index int, item mShipping.ShippingZoneCode) {
		zoneIds = append(zoneIds, item.ShippingZoneId)
	})
	zoneIds = slice.Unique(zoneIds)
	zoneListIn := ZoneListIn{ZoneIds: zoneIds}
	return s.ZoneList(zoneListIn)
}
