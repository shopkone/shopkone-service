package sShippingZoneFee

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/utility/handle"
)

// 更新规则
func (s *sShippingZoneFee) ConditionUpdate(fees []vo.BaseShippingZoneFee) error {

	feeIds := slice.Map(fees, func(index int, item vo.BaseShippingZoneFee) uint {
		return item.ID
	})
	feeIds = slice.Unique(feeIds)

	// 获取旧的条件
	var oldConditions []mShipping.ShippingZonFeeCondition
	if err := s.orm.Model(&oldConditions).
		Where("shop_id = ? AND shipping_zone_fee_id IN ?", s.shopId, feeIds).
		Omit("created_at", "deleted_at", "updated_at").Find(&oldConditions).Error; err != nil {
		return err
	}

	// 获取新的条件
	var newConditions []mShipping.ShippingZonFeeCondition
	for _, fee := range fees {
		for _, condition := range fee.Conditions {
			var temp mShipping.ShippingZonFeeCondition
			if err := convertor.CopyProperties(&temp, condition); err != nil {
				return err
			}
			temp.ShopId = s.shopId
			temp.ShippingZoneFeeId = fee.ID
			newConditions = append(newConditions, temp)
		}
	}

	// 获取差异
	insert, update, remove, err := handle.DiffUpdate(newConditions, oldConditions)
	if err != nil {
		return err
	}

	if len(insert) > 0 {
		var conditionCreateIn []ConditionCreateItem
		for _, condition := range insert {
			temp := ConditionCreateItem{}
			if err = convertor.CopyProperties(&temp, condition); err != nil {
				return err
			}
			conditionCreateIn = append(conditionCreateIn, temp)
		}
		if err = s.ConditionCreate(conditionCreateIn); err != nil {
			return err
		}
	}

	if len(remove) > 0 {
		removeIds := slice.Map(remove, func(index int, item mShipping.ShippingZonFeeCondition) uint {
			return item.ID
		})
		if err = s.orm.Model(&mShipping.ShippingZonFeeCondition{}).
			Where("id IN ? AND shop_id = ?", removeIds, s.shopId).
			Delete(&mShipping.ShippingZonFeeCondition{}).Error; err != nil {
			return err
		}
	}

	update = slice.Filter(update, func(index int, newCondition mShipping.ShippingZonFeeCondition) bool {
		oldCondition, ok := slice.FindBy(oldConditions, func(index int, item mShipping.ShippingZonFeeCondition) bool {
			return item.ID == newCondition.ID
		})
		return ok && s.ConditionIsChange(oldCondition, newCondition)
	})
	return s.ConditionUpdateBatch(update)
}
