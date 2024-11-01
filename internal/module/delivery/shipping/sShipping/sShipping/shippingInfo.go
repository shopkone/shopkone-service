package sShipping

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZone"
)

func (s *sShipping) ShippingInfo(id uint) (res vo.BaseShipping, err error) {
	// 获取物流方案
	var shipping mShipping.Shipping
	if err = s.orm.Model(&shipping).Where("shop_id = ? AND id = ?", s.shopId, id).
		Select("name", "type", "id").First(&shipping).Error; err != nil {
		return
	}
	res.ID = shipping.ID
	res.Name = shipping.Name
	res.Type = shipping.Type

	// 获取商品ids
	var shippingProducts []mShipping.ShippingProduct
	if err = s.orm.Model(&shippingProducts).Where("shop_id = ? AND shipping_id = ?", s.shopId, shipping.ID).
		Select("product_id").Find(&shippingProducts).Error; err != nil {
		return
	}
	res.ProductIDs = slice.Map(shippingProducts, func(_ int, item mShipping.ShippingProduct) uint {
		return item.ProductId
	})

	// 获取发货地点
	var shippingLocations []mShipping.ShippingLocation
	err = s.orm.Model(&shippingLocations).Where("shop_id = ? AND shipping_id = ?", s.shopId, shipping.ID).
		Select("location_id").Find(&shippingLocations).Error
	res.LocationIDs = slice.Map(shippingLocations, func(_ int, item mShipping.ShippingLocation) uint {
		return item.LocationId
	})

	// 获取发货区域
	res.Zones, err = sShippingZone.NewShippingZone(s.shopId, s.orm).ZoneList(shipping.ID)

	return res, err
}
