package sLocalDelivery

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
	"shopkone-service/internal/module/delivery/local-delivery/sLocalDelivery/sDeliveryArea"
)

func (s *sLocalDelivery) LocalDeliveryRemove(locationId uint) (err error) {
	// 获取本地配送
	var data mLocalDelivery.LocalDelivery
	if err = s.orm.Model(&data).Where("shop_id = ? AND location_id = ?", s.shopId, locationId).
		Select("id").First(&data).Error; err != nil {
		return err
	}

	// 删除本地配送
	if err = s.orm.Model(&mLocalDelivery.LocalDelivery{}).Where("shop_id = ? AND id = ?", s.shopId, data.ID).
		Delete(&mLocalDelivery.LocalDelivery{}).Error; err != nil {
		return err
	}

	// 删除配送区域
	sArea := sDeliveryArea.NewDeliveryArea(s.orm, s.shopId)
	areas, err := sArea.DeliveryAreasByDeliveryId(data.ID)
	areaIds := slice.Map(areas, func(index int, item mLocalDelivery.LocalDeliveryArea) uint {
		return item.ID
	})
	return sDeliveryArea.NewDeliveryArea(s.orm, s.shopId).DeliveryAreaRemove(areaIds)
}
