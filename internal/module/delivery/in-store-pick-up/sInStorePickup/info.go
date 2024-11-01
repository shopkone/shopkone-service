package sInStorePickup

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/in-store-pick-up/mInStorePickup"
)

func (s *sInStorePickup) Info(id uint) (info vo.InStorePickUpInfoRes, err error) {
	// 获取配送信息
	var data mInStorePickup.InStorePickup
	if err = s.orm.Model(&data).Where("shop_id = ? AND id = ?", s.shopId, id).
		Omit("shop_id", "created_at", "updated_at", "deleted_at").
		First(&data).Error; err != nil {
		return info, err
	}

	// 获取配送时间
	var hours []mInStorePickup.InStorePickupBusinessHours
	if err = s.orm.Model(&hours).Where("shop_id = ? AND in_store_pickup_id = ?", s.shopId, data.ID).
		Omit("shop_id", "created_at", "updated_at", "deleted_at").
		Find(&hours).Error; err != nil {
		return info, err
	}

	// 组装数据
	info.BaseInStorePickUp = vo.BaseInStorePickUp{
		Id:            data.ID,
		IsUnified:     data.IsUnified,
		Status:        data.Status,
		LocationID:    data.LocationId,
		PickupETA:     data.PickupETA,
		PickupETAUnit: data.PickupETAUnit,
		HasPickupETA:  data.HasPickupETA,
		End:           data.End,
		Start:         data.Start,
		Timezone:      data.Timezone,
		Weeks: slice.Map(hours, func(index int, item mInStorePickup.InStorePickupBusinessHours) vo.BaseInStorePickUpBusinessHours {
			i := vo.BaseInStorePickUpBusinessHours{}
			i.Id = item.ID
			i.Week = item.Week
			i.Start = item.Start
			i.End = item.End
			i.IsOpen = item.IsOpen
			return i
		}),
	}
	return info, err
}
