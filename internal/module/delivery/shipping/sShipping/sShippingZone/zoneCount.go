package sShippingZone

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

type CountByShippingIdsOut struct {
	Count      uint
	ShippingId uint
}

func (s *sShippingZone) ZoneCount(shippingIds []uint) (out []CountByShippingIdsOut, err error) {
	var zones []mShipping.ShippingZone
	if err = s.orm.Model(&zones).Where("shipping_id in (?) AND shop_id = ?", shippingIds, s.shopId).
		Select("id", "shipping_id").Find(&zones).Error; err != nil {
		return out, err
	}
	out = slice.Map(shippingIds, func(index int, item uint) CountByShippingIdsOut {
		i := CountByShippingIdsOut{}
		i.ShippingId = item
		currentZones := slice.Filter(zones, func(index int, i mShipping.ShippingZone) bool {
			return i.ShippingId == item
		})
		i.Count = uint(len(currentZones))
		return i
	})
	return out, err
}
