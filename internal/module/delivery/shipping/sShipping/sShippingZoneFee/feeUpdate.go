package sShippingZoneFee

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/utility/handle"
)

func (s *sShippingZoneFee) FeeUpdate(zones []vo.BaseShippingZone) error {
	zoneIds := slice.Map(zones, func(index int, item vo.BaseShippingZone) uint {
		return item.ID
	})
	zoneIds = slice.Unique(zoneIds)
	if len(zoneIds) == 0 {
		return nil
	}

	var newFees []mShipping.ShippingZoneFee
	for _, zone := range zones {
		for _, fee := range zone.Fees {
			temp := mShipping.ShippingZoneFee{}
			if err := convertor.CopyProperties(&temp, fee); err != nil {
				return err
			}
			temp.ShopId = s.shopId
			temp.ShippingZoneId = zone.ID
			newFees = append(newFees, temp)
		}
	}

	var oldFees []mShipping.ShippingZoneFee
	if err := s.orm.Model(mShipping.ShippingZoneFee{}).
		Where("shipping_zone_id IN ? AND shop_id = ?", zoneIds, s.shopId).
		Omit("created_at", "deleted_at", "updated_at").Find(&oldFees).Error; err != nil {
		return err
	}

	insert, update, remove, err := handle.DiffUpdate(newFees, oldFees)
	if err != nil {
		return err
	}

	// 新增
	if len(insert) > 0 {
		var feeCreateIn []FeeCreateItem
		slice.ForEach(zones, func(index int, zone vo.BaseShippingZone) {
			slice.ForEach(zone.Fees, func(index int, fee vo.BaseShippingZoneFee) {
				_, ok := slice.FindBy(insert, func(index int, item mShipping.ShippingZoneFee) bool {
					return item.ID == fee.ID
				})
				if ok {
					var feeCreateItem FeeCreateItem
					if err = convertor.CopyProperties(&feeCreateItem, fee); err != nil {
						return
					}
					feeCreateItem.ShippingZoneId = zone.ID
					feeCreateIn = append(feeCreateIn, feeCreateItem)
				}
			})
		})
		if err = s.FeeCreate(feeCreateIn); err != nil {
			return err
		}
	}

	// 更新
	if len(update) > 0 {
		var baseUpdateFees []vo.BaseShippingZoneFee
		slice.ForEach(zones, func(index int, zone vo.BaseShippingZone) {
			slice.ForEach(zone.Fees, func(index int, fee vo.BaseShippingZoneFee) {
				_, ok := slice.FindBy(update, func(index int, item mShipping.ShippingZoneFee) bool {
					return item.ID == fee.ID
				})
				if ok {
					baseUpdateFees = append(baseUpdateFees, fee)
				}
			})
		})
		if err = s.ConditionUpdate(baseUpdateFees); err != nil {
			return err
		}
		update = slice.Filter(update, func(index int, item mShipping.ShippingZoneFee) bool {
			find, ok := slice.FindBy(oldFees, func(index int, old mShipping.ShippingZoneFee) bool {
				return old.ID == item.ID
			})
			return ok && s.FeeIsChange(find, item)
		})
		if err = s.FeeUpdateBatch(update); err != nil {
			return err
		}
	}

	// 删除
	removeIds := slice.Map(remove, func(index int, item mShipping.ShippingZoneFee) uint {
		return item.ID
	})
	return s.FeeRemoveByIds(removeIds)
}
