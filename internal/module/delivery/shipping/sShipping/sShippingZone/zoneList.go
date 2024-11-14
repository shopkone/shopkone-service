package sShippingZone

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZoneFee"
)

type ZoneListIn struct {
	ShippingId uint
	ZoneIds    []uint
}

func (s *sShippingZone) ZoneList(in ZoneListIn) (res []vo.BaseShippingZone, err error) {
	sZoneFee := sShippingZoneFee.NewShippingZoneFee(s.orm, s.shopId)

	// 获取区域
	var zones []mShipping.ShippingZone
	query := s.orm.Model(&zones)
	if in.ShippingId != 0 {
		query = query.Where("shipping_id = ?", in.ShippingId)
	}
	if in.ZoneIds != nil && len(in.ZoneIds) > 0 {
		query = query.Where("id IN ?", in.ZoneIds)
	}
	if err = query.Select("id", "name").Find(&zones).Error; err != nil {
		return nil, err
	}
	zoneIds := slice.Map(zones, func(index int, item mShipping.ShippingZone) uint {
		return item.ID
	})

	// 获取codes
	var zoneCodes []mShipping.ShippingZoneCode
	if in.ShippingId != 0 {
		if err = s.orm.Model(&zoneCodes).Where("shipping_zone_id IN ? AND shop_id = ?", zoneIds, s.shopId).
			Select("shipping_zone_id", "country_code", "zone_codes", "id").Find(&zoneCodes).Error; err != nil {
			return nil, err
		}
	}

	// 获取运费
	var zoneFees []mShipping.ShippingZoneFee
	if err = s.orm.Model(&zoneFees).Where("shipping_zone_id IN ? AND shop_id = ?", zoneIds, s.shopId).
		Omit("shop_id", "updated_at", "created_at", "deleted_at").Find(&zoneFees).Error; err != nil {
		return nil, err
	}
	var zoneFeeIds []uint
	slice.ForEach(zoneFees, func(index int, item mShipping.ShippingZoneFee) {
		zoneFeeIds = append(zoneFeeIds, item.ID)
	})

	// 获取规则
	var zoneFeeRules []mShipping.ShippingZonFeeCondition
	if err = s.orm.Model(&zoneFeeRules).Where("shipping_zone_fee_id IN ?", zoneFeeIds).
		Omit("shop_id", "updated_at", "created_at", "deleted_at").Find(&zoneFeeRules).Error; err != nil {
		return nil, err
	}

	res = slice.Map(zones, func(index int, item mShipping.ShippingZone) vo.BaseShippingZone {
		i := vo.BaseShippingZone{}
		i.ID = item.ID
		i.Name = item.Name
		// 组装code
		codes := slice.Filter(zoneCodes, func(index int, code mShipping.ShippingZoneCode) bool {
			if code.ShippingZoneId == item.ID {
				return true
			}
			return false
		})
		i.Codes = slice.Map(codes, func(index int, item mShipping.ShippingZoneCode) vo.BaseZoneCode {
			return vo.BaseZoneCode{
				CountryCode: item.CountryCode,
				ZoneCodes:   item.ZoneCodes,
				Id:          item.ID,
			}
		})
		// 组装费用
		fees := slice.Filter(zoneFees, func(index int, fee mShipping.ShippingZoneFee) bool {
			return fee.ShippingZoneId == item.ID
		})
		i.Fees = slice.Map(fees, func(index int, item mShipping.ShippingZoneFee) vo.BaseShippingZoneFee {
			temp := sZoneFee.ModelToFee(item)
			// 组装规则
			conditions := slice.Filter(zoneFeeRules, func(index int, rule mShipping.ShippingZonFeeCondition) bool {
				return rule.ShippingZoneFeeId == item.ID
			})
			temp.Conditions = slice.Map(conditions, func(index int, item mShipping.ShippingZonFeeCondition) vo.BaseShippingZoneFeeCondition {
				return sZoneFee.ModelToCondition(item)
			})
			return temp
		})
		return i
	})

	return res, err
}
