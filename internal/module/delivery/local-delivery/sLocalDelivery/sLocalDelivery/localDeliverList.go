package sLocalDelivery

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
)

func (s *sLocalDelivery) LocalDeliveryList(locationIds []uint) (res []vo.LocalDeliveryListRes, err error) {
	var list []mLocalDelivery.LocalDelivery
	if err = s.orm.Model(&list).Where("location_id IN ?", locationIds).
		Where("shop_id = ?", s.shopId).
		Omit("updated_at", "created_at", "deleted_at", "shop_id").
		Find(&list).Error; err != nil {
		return res, err
	}
	return slice.Map(list, func(index int, item mLocalDelivery.LocalDelivery) vo.LocalDeliveryListRes {
		return vo.LocalDeliveryListRes{
			Id:         item.ID,
			Status:     item.Status,
			LocationId: item.LocationId,
		}
	}), nil
}
