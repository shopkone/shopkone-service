package sInStorePickup

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/in-store-pick-up/mInStorePickup"
)

func (s *sInStorePickup) List(locationIds []uint) (res []vo.InStorePickUpListRes, err error) {
	var data []mInStorePickup.InStorePickup
	if err = s.orm.Model(&mInStorePickup.InStorePickup{}).
		Omit("shop_id", "deleted_at", "updated_at", "created_at").
		Where("shop_id = ? AND location_id IN ?", s.shopId, locationIds).Find(&data).Error; err != nil {
		return res, err
	}

	res = slice.Map(data, func(index int, item mInStorePickup.InStorePickup) vo.InStorePickUpListRes {
		i := vo.InStorePickUpListRes{}
		i.Id = item.ID
		i.Status = item.Status
		i.LocationID = item.LocationId
		return i
	})

	return res, nil
}
