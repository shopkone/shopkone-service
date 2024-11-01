package sShipping

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

func (s *sShipping) LocationUpdate(shippingId uint, newLocationIds []uint) (err error) {
	// 获取旧的locationIds
	var oldLocations []mShipping.ShippingLocation
	if err = s.orm.Model(&oldLocations).Where("shipping_id = ? AND shop_id = ?", shippingId, s.shopId).
		Select("location_id").Find(&oldLocations).Error; err != nil {
		return err
	}
	oldLocationIds := slice.Map(oldLocations, func(_ int, item mShipping.ShippingLocation) uint {
		return item.LocationId
	})

	// 找出要删除的locationIds
	deleteLocations := slice.Filter(oldLocationIds, func(_ int, item uint) bool {
		return slice.Contain(newLocationIds, item) == false
	})
	if err = s.LocationRemove(shippingId, deleteLocations); err != nil {
		return err
	}

	// 添加新的
	addLocationIds := slice.Filter(newLocationIds, func(_ int, item uint) bool {
		return slice.Contain(oldLocationIds, item) == false
	})
	return s.LocationCreate(shippingId, addLocationIds)
}
