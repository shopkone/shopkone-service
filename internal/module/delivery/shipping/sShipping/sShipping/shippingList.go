package sShipping

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZone"
)

func (s *sShipping) ShippingList() (res []vo.ShippingListRes, err error) {
	// 获取物流列表
	var shippings []mShipping.Shipping
	if err = s.orm.Model(&shippings).Where("shop_id = ?", s.shopId).
		Select("id", "type", "name").Find(&shippings).Error; err != nil {
		return res, err
	}
	shippingIds := slice.Map(shippings, func(index int, item mShipping.Shipping) uint {
		return item.ID
	})

	// 获取物流商品
	var shippingProducts []mShipping.ShippingProduct
	if err = s.orm.Model(&shippingProducts).Where("shipping_id in (?) AND shop_id = ?", shippingIds, s.shopId).
		Select("shipping_id").Find(&shippingProducts).Error; err != nil {
		return res, err
	}

	// 获取发货地点
	var shippingLocations []mShipping.ShippingLocation
	if err = s.orm.Model(&shippingLocations).Where("shipping_id in (?)", shippingIds).
		Select("shipping_id").Find(&shippingLocations).Error; err != nil {
		return res, err
	}

	// 获取发货区域
	zoneCounts, err := sShippingZone.NewShippingZone(s.shopId, s.orm).ZoneCount(shippingIds)
	if err != nil {
		return nil, err
	}

	// 组装数据
	res = slice.Map(shippings, func(index int, item mShipping.Shipping) vo.ShippingListRes {
		i := vo.ShippingListRes{}
		i.Name = item.Name
		i.Type = item.Type
		i.Id = item.ID
		i.ProductCount = uint(slice.CountBy(shippingProducts, func(index int, z mShipping.ShippingProduct) bool {
			return z.ShippingId == item.ID
		}))
		i.LocationCount = uint(slice.CountBy(shippingLocations, func(index int, z mShipping.ShippingLocation) bool {
			return z.ShippingId == item.ID
		}))
		zone, ok := slice.FindBy(zoneCounts, func(index int, z sShippingZone.CountByShippingIdsOut) bool {
			return z.ShippingId == item.ID
		})
		if ok {
			i.ZoneCount = zone.Count
		}
		return i
	})

	return res, err
}
