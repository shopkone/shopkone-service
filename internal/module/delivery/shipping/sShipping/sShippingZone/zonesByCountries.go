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
	res, err = s.ZoneList(zoneListIn)
	// 从上面的数据中，获取国家
	res = slice.Map(res, func(index int, item vo.BaseShippingZone) vo.BaseShippingZone {
		zonesCodes := slice.Filter(zoneCodes, func(index int, i mShipping.ShippingZoneCode) bool {
			return i.ShippingZoneId == item.ID
		})
		if len(zonesCodes) > 0 {
			item.Codes = slice.Map(zonesCodes, func(index int, z mShipping.ShippingZoneCode) vo.BaseZoneCode {
				// 不获取区域，只要国家
				return vo.BaseZoneCode{
					CountryCode: z.CountryCode,
					Id:          item.ID,
				}
			})
		}
		return item
	})
	return res, err
}
