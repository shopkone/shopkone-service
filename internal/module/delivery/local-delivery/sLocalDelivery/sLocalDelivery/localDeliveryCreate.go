package sLocalDelivery

import (
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
)

func (s *sLocalDelivery) LocalDeliveryCreate(locationId uint) (err error) {
	data := mLocalDelivery.LocalDelivery{}
	data.Status = mLocalDelivery.LocalDeliveryStatusClose
	data.LocationId = locationId
	data.ShopId = s.shopId
	return s.orm.Create(&data).Error
}
