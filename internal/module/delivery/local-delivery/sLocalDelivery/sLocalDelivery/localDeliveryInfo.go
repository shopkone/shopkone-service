package sLocalDelivery

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
	"shopkone-service/internal/module/delivery/local-delivery/sLocalDelivery/sDeliveryArea"
)

func (s *sLocalDelivery) LocalDeliveryInfo(id uint) (res vo.LocalDeliveryInfoRes, err error) {
	// 获取本地配送
	var localDelivery mLocalDelivery.LocalDelivery
	if err = s.orm.Model(&localDelivery).Where("shop_id = ? AND id = ?", s.shopId, id).
		Omit("updated_at", "created_at", "deleted_at", "shop_id").
		First(&localDelivery).Error; err != nil {
		return res, err
	}

	res.Id = localDelivery.ID
	res.Status = localDelivery.Status
	res.LocationId = localDelivery.LocationId

	res.Areas, err = sDeliveryArea.NewDeliveryArea(s.orm, s.shopId).DeliveryAreaList(id)
	return res, err
}
