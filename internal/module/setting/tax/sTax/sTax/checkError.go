package sTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/resource/mResource"
	"shopkone-service/internal/module/base/resource/sResource"
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/utility/code"
)

func (s *sTax) CheckError(taxId uint) (info vo.TaxInfoRes, err error) {
	info, err = s.TaxInfo(taxId)
	if err != nil {
		return info, err
	}

	country, err := sResource.NewCountry().CountryByCode(info.CountryCode)
	if err != nil {
		return info, err
	}

	// 如果有区域税率，校验区域税率
	if len(info.Zones) > 0 {
		// 区域是否存在
		isAllExist := slice.Every(info.Zones, func(index int, zone vo.BaseTaxZone) bool {
			_, has := slice.FindBy(country.Zones, func(index int, item mResource.CountryZone) bool {
				return item.Code == zone.ZoneCode
			})
			return has
		})
		if !isAllExist {
			return info, code.TaxZoneNotFound
		}

		// 区域是否重复
		zoneCodes := slice.Map(info.Zones, func(index int, zone vo.BaseTaxZone) string {
			return zone.ZoneCode
		})
		if len(slice.Unique(zoneCodes)) != len(info.Zones) {
			return info, code.TaxZooeCodeRepeat
		}

		// 税率是否小于0
		isAllThenZero := slice.Every(info.Zones, func(index int, zone vo.BaseTaxZone) bool {
			return zone.TaxRate >= 0
		})
		if !isAllThenZero {
			return info, code.TaxRateLessThanZero
		}
	}

	// 如果有自定税率，校验自定义税率
	if len(info.Customers) > 0 {
		// 是否都有zones
		isAllHasZones := slice.Every(info.Customers, func(index int, customer vo.BaseCustomerTax) bool {
			return len(customer.Zones) > 0
		})
		if !isAllHasZones {
			return info, code.TaxZonesMust
		}

		// 如果是系列类型，则必须要有系列
		isAllHasCollection := slice.Every(info.Customers, func(index int, customer vo.BaseCustomerTax) bool {
			if customer.Type == mTax.CustomerTaxTypeCollection {
				return customer.CollectionID > 0
			}
			return true
		})
		if !isAllHasCollection {
			return info, code.TaxCollectionMust
		}

		// collectionId 是否全部都是存在的 TODO:在api实现了，原有是依赖循环，待处理

		// collectionId 是否重复了
		collectionIds := slice.Map(info.Customers, func(index int, customer vo.BaseCustomerTax) uint {
			return customer.CollectionID
		})
		if len(slice.Unique(collectionIds)) != len(collectionIds) {
			return info, code.TaxCollectionRepeat
		}

		// 区域是否存在
		isAllExist := slice.Every(info.Customers, func(index int, customer vo.BaseCustomerTax) bool {
			return slice.Every(customer.Zones, func(index int, zone vo.BaseCustomerTaxZone) bool {
				_, has := slice.FindBy(country.Zones, func(index int, item mResource.CountryZone) bool {
					return item.Code == zone.AreaCode
				})
				return has || zone.AreaCode == country.Code
			})
		})
		if !isAllExist {
			return info, code.TaxZoneNotFound
		}

		// 区域是否重复
		isAllNoRepeat := slice.Every(info.Customers, func(index int, customer vo.BaseCustomerTax) bool {
			codes := slice.Map(customer.Zones, func(index int, item vo.BaseCustomerTaxZone) string {
				return item.AreaCode
			})
			return len(customer.Zones) == len(slice.Unique(codes))
		})
		if !isAllNoRepeat {
			return info, code.TaxZooeCodeRepeat
		}

		// 税率是否小于0
		isAllThenZero := slice.Every(info.Customers, func(index int, customer vo.BaseCustomerTax) bool {
			return slice.Every(customer.Zones, func(index int, zone vo.BaseCustomerTaxZone) bool {
				return zone.TaxRate >= 0
			})
		})
		if !isAllThenZero {
			return info, code.TaxRateLessThanZero
		}
	}

	return info, err
}
