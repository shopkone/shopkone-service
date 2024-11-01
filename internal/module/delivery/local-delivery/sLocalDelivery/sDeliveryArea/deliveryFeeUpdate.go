package sDeliveryArea

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
	"shopkone-service/utility/handle"
)

func (s *sDeliveryArea) DeliveryFeeUpdate(newFees []mLocalDelivery.LocalDeliveryFee) error {
	if len(newFees) == 0 {
		return nil
	}

	var areaIds []uint
	for _, item := range newFees {
		areaIds = append(areaIds, item.LocalDeliveryAreaID)
	}
	areaIds = slice.Unique(areaIds)

	// 获取旧的配送费用
	var oldFees []mLocalDelivery.LocalDeliveryFee
	if err := s.orm.Model(&oldFees).Where("local_delivery_area_id in ?", areaIds).
		Omit("created_at", "updated_at", "deleted_at").
		Where("shop_id = ?", s.shopId).Find(&oldFees).Error; err != nil {
		return err
	}

	// 变更
	insert, update, remove, err := handle.DiffUpdate(newFees, oldFees)
	if err != nil {
		return err
	}

	if len(insert) > 0 {
		insert = slice.Map(insert, func(index int, item mLocalDelivery.LocalDeliveryFee) mLocalDelivery.LocalDeliveryFee {
			item.ID = 0
			return item
		})
		if err = s.orm.Create(insert).Error; err != nil {
			return err
		}
	}

	if len(update) > 0 {
		update = slice.Filter(update, func(index int, item mLocalDelivery.LocalDeliveryFee) bool {
			old, ok := slice.FindBy(oldFees, func(index int, old mLocalDelivery.LocalDeliveryFee) bool {
				return old.ID == item.ID
			})
			if !ok {
				return false
			}
			return s.DeliveryFeeIsChange(old, item)
		})
		if err = s.DeliveryFeeUpdateBatch(update); err != nil {
			return err
		}
	}

	if len(remove) > 0 {
		return s.orm.Delete(&remove).Error
	}

	return err
}
