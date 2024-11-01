package sShipping

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/resource/mResource"
	"shopkone-service/internal/module/base/resource/sResource"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/product/product/sProduct/sProduct"
	"shopkone-service/internal/module/setting/location/sLocation"
	"shopkone-service/utility/code"
)

// TODO:
func (s *sShipping) Check(in vo.BaseShipping, id uint) (err error) {
	// 先讨论通用运费的情况
	if in.Type == mShipping.GeneralExpressDelivery && id == 0 {
		// 通用方案只允许有一个
		if err = s.CheckGeneralExist(); err != nil {
			return err
		}
	}

	// 再讨论自定义运费的情况
	if in.Type == mShipping.CustomerExpressDelivery {
		// 物流方案名称不能为空
		if in.Name == "" {
			return code.ShippingNameMust
		}
		// 物流方案名称不能重复
		if err = s.CheckNameRepeat(in.Name, id); err != nil {
			return err
		}
		//发货地点必须
		if len(in.LocationIDs) == 0 {
			return code.LocationMust
		}
		// 发货地点必须可用
		if err = sLocation.NewLocation(s.orm, s.shopId).IsAllActive(in.LocationIDs); err != nil {
			return err
		}
		// 商品必须存在
		if len(in.ProductIDs) > 0 {
			if err = sProduct.NewProduct(s.orm, s.shopId).CheckExist(in.ProductIDs); err != nil {
				return err
			}
		}
	}

	// 区域名称不能重复、为空
	if len(in.Zones) == 0 {
		return code.ZoneMust
	}
	zoneNames := slice.Map(in.Zones, func(index int, item vo.BaseShippingZone) string {
		return item.Name
	})
	if someEmpty := slice.Some(zoneNames, func(index int, item string) bool {
		return item == ""
	}); someEmpty {
		return code.ShippingZoneMust
	}
	if len(slice.Unique(zoneNames)) != len(zoneNames) {
		return code.ShippingZoneNameRepeat
	}

	// 单个区域内的运费名称不能重复、为空
	if someEmpty := slice.Some(in.Zones, func(index int, zone vo.BaseShippingZone) bool {
		return slice.Some(zone.Fees, func(index int, fee vo.BaseShippingZoneFee) bool {
			return fee.Name == ""
		}) || len(zone.Fees) == 0
	}); someEmpty {
		return code.ShippingFeeNameMust
	}
	if someRepeat := slice.Some(in.Zones, func(index int, zone vo.BaseShippingZone) bool {
		names := slice.Map(zone.Fees, func(index int, fee vo.BaseShippingZoneFee) string {
			return fee.Name
		})
		return len(slice.Unique(names)) != len(names)
	}); someRepeat {
		return code.ZoneMust
	}

	// 确保所有的区域码都有效
	sCountry := sResource.NewCountry()
	isEmpty := slice.Some(in.Zones, func(index int, zone vo.BaseShippingZone) bool {
		if len(zone.Codes) == 0 {
			return true
		}
		return slice.Some(zone.Codes, func(index int, code vo.BaseZoneCode) bool {
			country, err := sCountry.CountryByCode(code.CountryCode)
			if err != nil {
				return true
			}
			return slice.Some(code.ZoneCodes, func(index int, zoneCode string) bool {
				_, ok := slice.FindBy(country.Zones, func(index int, i mResource.CountryZone) bool {
					return i.Code == zoneCode
				})
				return !ok
			})
		})
	})

	if err != nil {
		return err
	}
	if isEmpty {
		return code.ZoneCodeMust
	}
	return err
}
