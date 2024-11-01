package sInStorePickup

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/in-store-pick-up/mInStorePickup"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

func (s *sInStorePickup) Update(in vo.InStorePickUpUpdateReq) error {
	// 获取info
	var info mInStorePickup.InStorePickup
	if err := s.orm.Model(&info).Where("id = ? AND shop_id = ?", in.Id, s.shopId).
		Select("id", "status").First(&info).Error; err != nil {
		return err
	}
	info.Start = in.Start
	info.End = in.End
	info.PickupETA = in.PickupETA
	info.PickupETAUnit = in.PickupETAUnit
	info.HasPickupETA = in.HasPickupETA
	info.IsUnified = in.IsUnified
	info.Status = in.Status
	info.Timezone = in.Timezone

	// 更新
	if err := s.orm.Model(&info).Where("id = ? AND shop_id = ?", in.Id, s.shopId).
		Select("status", "start", "end", "pickup_eta", "pickup_eta_unit", "has_pickup_eta", "is_unified", "timezone").
		Updates(&info).Error; err != nil {
		return err
	}

	// 更新weeks
	if len(in.Weeks) != 7 {
		return code.IdMissing
	}
	newWeeks := slice.Map(in.Weeks, func(index int, item vo.BaseInStorePickUpBusinessHours) mInStorePickup.InStorePickupBusinessHours {
		i := mInStorePickup.InStorePickupBusinessHours{}
		i.ID = item.Id
		i.Week = item.Week
		i.Start = item.Start
		i.End = item.End
		i.IsOpen = item.IsOpen
		return i
	})
	// 获取旧的weeks
	var oldWeeks []mInStorePickup.InStorePickupBusinessHours
	if err := s.orm.Model(&oldWeeks).Where("shop_id = ? AND in_store_pickup_id = ?", s.shopId, info.ID).
		Omit("updated_at", "created_at", "deleted_at").Find(&oldWeeks).Error; err != nil {
		return err
	}

	// 获取差异，不允许新增或者删除
	insert, update, remove, err := handle.DiffUpdate(newWeeks, oldWeeks)
	if err != nil {
		return err
	}
	if len(insert) != 0 || len(remove) != 0 {
		return code.IdMissing
	}

	update = slice.Filter(update, func(index int, newItem mInStorePickup.InStorePickupBusinessHours) bool {
		oldItem, ok := slice.FindBy(oldWeeks, func(index int, item mInStorePickup.InStorePickupBusinessHours) bool {
			return item.Week == newItem.Week
		})
		if !ok {
			return false
		}
		return s.WeekIsChange(oldItem, newItem)
	})
	if len(update) == 0 {
		return err
	}
	update = slice.Map(update, func(index int, item mInStorePickup.InStorePickupBusinessHours) mInStorePickup.InStorePickupBusinessHours {
		item.CanCreateId = true
		item.ShopId = s.shopId
		item.InStorePickupID = info.ID
		return item
	})
	batchIn := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"is_open", "start", "end"},
	}
	if err = handle.BatchUpdateById(batchIn, &update); err != nil {
		return err
	}
	return err
}
