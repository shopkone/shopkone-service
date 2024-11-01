package sShippingZone

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

func (s *sShippingZone) CodesUpdate(zones []vo.BaseShippingZone) error {
	zoneIds := slice.Map(zones, func(index int, item vo.BaseShippingZone) uint {
		return item.ID
	})

	// 新的
	var newCodes []mShipping.ShippingZoneCode
	slice.ForEach(zones, func(index int, zone vo.BaseShippingZone) {
		slice.ForEach(zone.Codes, func(index int, code vo.BaseZoneCode) {
			temp := mShipping.ShippingZoneCode{}
			temp.CountryCode = code.CountryCode
			temp.ZoneCodes = code.ZoneCodes
			temp.ShopId = s.shopId
			temp.ShippingZoneId = zone.ID
			newCodes = append(newCodes, temp)
		})
	})

	// 旧的
	var oldNewCodes []mShipping.ShippingZoneCode
	if err := s.orm.Model(&mShipping.ShippingZoneCode{}).Where("shop_id = ? AND shipping_zone_id IN ?", s.shopId, zoneIds).
		Omit("deleted_at", "created_at", "updated_at").Find(&oldNewCodes).Error; err != nil {
		return err
	}

	// 是否有变更
	change := s.CodesIsChange(newCodes, oldNewCodes)
	removeIds := slice.Map(change.Remove, func(index int, item mShipping.ShippingZoneCode) uint {
		return item.ID
	})

	if len(removeIds) > 0 {
		if err := s.orm.Model(&mShipping.ShippingZoneCode{}).Where("id IN ? AND shop_id = ?", removeIds, s.shopId).
			Unscoped().Delete(&mShipping.ShippingZoneCode{}).Error; err != nil {
			return err
		}
	}
	if len(change.Insert) > 0 {
		// 创建国家
		if err := s.CodesCreate(change.Insert); err != nil {
			return err
		}
	}

	return nil
}
