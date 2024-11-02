package sShippingZone

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZoneFee"
	"shopkone-service/utility/handle"
)

// 更新区域
func (s *sShippingZone) ZoneUpdate(zones []vo.BaseShippingZone, shippingId uint) error {
	// 获取旧的区域
	var oldZones []mShipping.ShippingZone
	if err := s.orm.Model(&oldZones).Where("shipping_id = ? AND shop_id = ?", shippingId, s.shopId).
		Select("id", "name").Find(&oldZones).Error; err != nil {
		return err
	}

	// 获取新的区域
	newZones := slice.Map(zones, func(index int, item vo.BaseShippingZone) mShipping.ShippingZone {
		i := mShipping.ShippingZone{}
		i.ID = item.ID
		i.ShopId = s.shopId
		i.ShippingId = shippingId
		i.Name = item.Name
		return i
	})

	// 找出差异
	insert, update, remove, err := handle.DiffUpdate(newZones, oldZones)
	if err != nil {
		return err
	}

	// 插入区域
	baseInsert := slice.Filter(zones, func(index int, zone vo.BaseShippingZone) bool {
		_, ok := slice.FindBy(insert, func(index int, item mShipping.ShippingZone) bool {
			return item.ID == zone.ID
		})
		return ok
	})
	if len(baseInsert) > 0 {
		if err = s.ZoneCreate(baseInsert, shippingId); err != nil {
			return err
		}
	}

	// 更新区域费用
	baseUpdate := slice.Filter(zones, func(index int, zone vo.BaseShippingZone) bool {
		_, ok := slice.FindBy(update, func(index int, item mShipping.ShippingZone) bool {
			return item.ID == zone.ID
		})
		return ok
	})
	if err = sShippingZoneFee.NewShippingZoneFee(s.orm, s.shopId).FeeUpdate(baseUpdate); err != nil {
		return err
	}
	// 更新区域码
	if err = s.CodesUpdate(baseUpdate); err != nil {
		return err
	}
	// 更新区域
	update = slice.Filter(update, func(index int, newZone mShipping.ShippingZone) bool {
		oldZone, ok := slice.FindBy(oldZones, func(index int, oldZone mShipping.ShippingZone) bool {
			return oldZone.ID == newZone.ID
		})
		return ok && newZone.Name != oldZone.Name
	})
	if err = s.ZoneUpdateBatch(update); err != nil {
		return err
	}

	// 删除区域
	removeIds := slice.Map(remove, func(index int, item mShipping.ShippingZone) uint {
		return item.ID
	})
	if err = s.ZoneRemove(removeIds); err != nil {
		return err
	}

	return err
}
