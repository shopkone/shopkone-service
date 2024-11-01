package sInStorePickup

import (
	"shopkone-service/internal/module/delivery/in-store-pick-up/mInStorePickup"
	"shopkone-service/utility/handle"
	"time"
)

func (s *sInStorePickup) Create(locationId uint, timezone string) error {
	sixty := uint(60)

	// 自提
	var data mInStorePickup.InStorePickup
	data.Status = mInStorePickup.InStorePickupStatusClose
	data.IsUnified = true
	data.LocationId = locationId
	data.ShopId = s.shopId
	data.PickupETA = &sixty
	data.PickupETAUnit = mInStorePickup.InStorePickupTimeUnitMinute
	data.HasPickupETA = true
	data.Start = 9 * 60
	data.End = 17 * 60
	data.Timezone = timezone
	if err := s.orm.Create(&data).Error; err != nil {
		return err
	}

	// 自提点时间
	var hours []mInStorePickup.InStorePickupBusinessHours
	for i := 0; i < 7; i++ {
		hour := mInStorePickup.InStorePickupBusinessHours{}
		hour.Week = uint8(i)
		hour.IsOpen = true
		hour.Start = 9 * 60
		hour.End = 17 * 60
		hour.InStorePickupID = data.ID
		hour.ShopId = data.ShopId
		hours = append(hours, hour)
	}
	return s.orm.Create(&hours).Error
}

func GetHourTime(hour int) time.Time {
	t := time.Date(2006, 1, 1, hour, 0, 0, 0, time.UTC)
	return handle.TranslateTime(t)
}
