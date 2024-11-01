package sInStorePickup

import "shopkone-service/internal/module/delivery/in-store-pick-up/mInStorePickup"

func (s *sInStorePickup) WeekIsChange(oldWeek, newWeek mInStorePickup.InStorePickupBusinessHours) bool {
	if oldWeek.Start != newWeek.Start {
		return true
	}
	if oldWeek.End != newWeek.End {
		return true
	}
	if oldWeek.IsOpen != newWeek.IsOpen {
		return true
	}
	return false
}
