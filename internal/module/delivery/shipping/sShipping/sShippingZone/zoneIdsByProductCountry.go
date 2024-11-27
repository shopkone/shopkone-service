package sShippingZone

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

type ZoneIdsByCountryIn struct {
	CountryCode string
	ZoneCode    string
}

type ZoneIdsByCountryOut struct {
	ShippingId uint
	ZoneIds    []uint
}

func (s *sShippingZone) ZoneIdsByCountry(in ZoneIdsByCountryIn) (out []ZoneIdsByCountryOut, err error) {
	// 根据国家获取区域码
	var zoneCodes []mShipping.ShippingZoneCode
	if err = s.orm.Where("country_code = ?", in.CountryCode).
		Omit("updated_at", "created_at", "deleted_at").
		Find(&zoneCodes).Error; err != nil {
		return out, err
	}
	// 筛选出符合的区域
	if in.ZoneCode != "" && len(zoneCodes) > 0 {
		zoneCodes = slice.Filter(zoneCodes, func(index int, item mShipping.ShippingZoneCode) bool {
			return slice.Contain(item.ZoneCodes, in.ZoneCode)
		})
	}
	// 获取区域码的区域
	zoneIds := slice.Map(zoneCodes, func(index int, item mShipping.ShippingZoneCode) uint {
		return item.ShippingZoneId
	})
	zoneIds = slice.Unique(zoneIds)
	// 获取区域
	var zones []mShipping.ShippingZone
	if err = s.orm.Where("id in ?", zoneIds).
		Omit("updated_at", "created_at", "deleted_at").
		Find(&zones).Error; err != nil {
		return out, err
	}
	slice.ForEach(zones, func(index int, zone mShipping.ShippingZone) {
		o, ok := slice.FindBy(out, func(index int, o ZoneIdsByCountryOut) bool {
			return zone.ShippingId == o.ShippingId
		})
		if ok {
			o.ZoneIds = append(o.ZoneIds, zone.ID)
		} else {
			out = append(out, ZoneIdsByCountryOut{
				ShippingId: zone.ShippingId,
				ZoneIds:    []uint{zone.ID},
			})
		}
	})
	return out, err
}
