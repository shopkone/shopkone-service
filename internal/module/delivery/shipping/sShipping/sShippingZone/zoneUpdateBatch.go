package sShippingZone

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/utility/handle"
)

type ZoneUpdateItem struct{}

func (s *sShippingZone) ZoneUpdateBatch(zones []mShipping.ShippingZone) error {
	if len(zones) == 0 {
		return nil
	}
	// 更新区域
	zones = slice.Map(zones, func(index int, item mShipping.ShippingZone) mShipping.ShippingZone {
		item.CanCreateId = true
		return item
	})
	in := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"name"},
	}
	return handle.BatchUpdateById(in, &zones)
}
