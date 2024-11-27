package sShipping

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZone"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZoneFee"
)

type ShippingFeesByCountryProductIn struct {
	CountryCode string
	ZoneCode    string
	ProductIDs  []uint
}
type ShippingFeesByCountryProductOut struct {
}

func (s *sShipping) FeesByCountryProduct(in ShippingFeesByCountryProductIn) (fees []mShipping.ShippingZoneFee, err error) {
	// 获取区域
	zoneService := sShippingZone.NewShippingZone(s.shopId, s.orm)

	zonesByCountryIn := sShippingZone.ZoneIdsByCountryIn{
		CountryCode: in.CountryCode,
		ZoneCode:    in.ZoneCode,
	}
	zones, err := zoneService.ZoneIdsByCountry(zonesByCountryIn)
	if err != nil {
		return fees, err
	}
	shippingIds := slice.Map(zones, func(index int, item sShippingZone.ZoneIdsByCountryOut) uint {
		return item.ShippingId
	})

	// 找出来包含这些商品的
	var shippingProducts []mShipping.ShippingProduct
	if err = s.orm.Model(&shippingProducts).Where("product_id in (?)", in.ProductIDs).
		Where("product_id IN (?)", in.ProductIDs).
		Select("shipping_id", "product_id").
		Find(&shippingProducts).Error; err != nil {
		return fees, err
	}
	canUseShippingIds := slice.Filter(shippingIds, func(index int, id uint) bool {
		_, has := slice.FindBy(shippingProducts, func(index int, item mShipping.ShippingProduct) bool {
			return item.ShippingId == id
		})
		return has
	})
	// 商品需要全部匹配才可以使用
	canUseShippingIds = slice.Filter(canUseShippingIds, func(index int, id uint) bool {
		ps := slice.Filter(shippingProducts, func(index int, item mShipping.ShippingProduct) bool {
			return item.ShippingId == id
		})
		isProductAllMatch := slice.Every(in.ProductIDs, func(index int, productId uint) bool {
			_, ok := slice.FindBy(ps, func(index int, item mShipping.ShippingProduct) bool {
				return item.ProductId == productId
			})
			return ok
		})
		return isProductAllMatch
	})

	genderShipping, err := s.ShippingGeneral()
	if err != nil {
		return fees, err
	}

	// 筛选可用区域
	canZones := slice.Filter(zones, func(index int, item sShippingZone.ZoneIdsByCountryOut) bool {
		if item.ShippingId == 0 {
			return false
		}
		if item.ShippingId == genderShipping.ID {
			return true
		}
		_, ok := slice.FindBy(canUseShippingIds, func(index int, id uint) bool {
			return item.ShippingId == id
		})
		return ok
	})

	// 根据可用区域获取费用
	var zoneIds []uint
	slice.ForEach(canZones, func(index int, item sShippingZone.ZoneIdsByCountryOut) {
		zoneIds = append(zoneIds, item.ZoneIds...)
	})
	// 获取费用
	zoneFeeService := sShippingZoneFee.NewShippingZoneFee(s.orm, s.shopId)

	return zoneFeeService.FeesByZoneIds(zoneIds)
}
