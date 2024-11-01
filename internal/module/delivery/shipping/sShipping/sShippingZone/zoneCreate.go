package sShippingZone

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZoneFee"
	"shopkone-service/utility/code"
)

func (s *sShippingZone) ZoneCreate(zones []vo.BaseShippingZone, shippingId uint) (err error) {
	if len(zones) == 0 {
		return code.ZoneMust
	}

	sZoneFee := sShippingZoneFee.NewShippingZoneFee(s.orm, s.shopId)

	// 创建区域
	shippingZones := slice.Map(zones, func(index int, item vo.BaseShippingZone) mShipping.ShippingZone {
		i := mShipping.ShippingZone{}
		i.Name = item.Name
		i.ShippingId = shippingId
		i.ShopId = s.shopId
		return i
	})
	if err = s.orm.Create(&shippingZones).Error; err != nil {
		return err
	}

	var zoneCodes []mShipping.ShippingZoneCode
	var createFeeListIn []sShippingZoneFee.FeeCreateItem
	for _, zone := range zones {
		realZone, ok := slice.FindBy(shippingZones, func(index int, real mShipping.ShippingZone) bool {
			return real.Name == zone.Name
		})
		if !ok {
			return code.ZoneMust
		}
		// 区域国家
		if len(zone.Codes) == 0 {
			return code.ZoneCodeMust
		}
		slice.ForEach(zone.Codes, func(index int, item vo.BaseZoneCode) {
			zoneCode := mShipping.ShippingZoneCode{}
			zoneCode.CountryCode = item.CountryCode
			zoneCode.ZoneCodes = item.ZoneCodes
			zoneCode.ShopId = s.shopId
			zoneCode.ShippingZoneId = realZone.ID
			zoneCodes = append(zoneCodes, zoneCode)
		})

		// 区域费用
		if len(zone.Fees) == 0 {
			return code.ZoneFeeMust
		}
		for _, fee := range zone.Fees {
			var createFeeIn sShippingZoneFee.FeeCreateItem
			if err = convertor.CopyProperties(&createFeeIn, fee); err != nil {
				return
			}
			createFeeIn.ShippingZoneId = realZone.ID
			createFeeListIn = append(createFeeListIn, createFeeIn)
		}
	}

	// 创建区域国家
	if err = s.CodesCreate(zoneCodes); err != nil {
		return err
	}

	// 创建区域运费方案
	return sZoneFee.FeeCreate(createFeeListIn)
}
