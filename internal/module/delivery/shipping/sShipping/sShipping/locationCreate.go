package sShipping

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

func (s *sShipping) LocationCreate(shippingId uint, locationIds []uint) (err error) {
	if len(locationIds) == 0 {
		return nil
	}
	locations := slice.Map(locationIds, func(index int, item uint) mShipping.ShippingLocation {
		i := mShipping.ShippingLocation{}
		i.ShippingId = shippingId
		i.LocationId = item
		i.ShopId = s.shopId
		return i
	})
	return s.orm.Create(&locations).Error
}
