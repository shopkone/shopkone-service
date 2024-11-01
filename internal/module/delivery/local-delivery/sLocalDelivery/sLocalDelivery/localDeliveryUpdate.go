package sLocalDelivery

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/local-delivery/mLocalDelivery"
	"shopkone-service/internal/module/delivery/local-delivery/sLocalDelivery/sDeliveryArea"
)

func (s *sLocalDelivery) LocalDeliveryUpdate(in vo.UpdateLocalDeliveryReq) (err error) {
	// 更新本地配送
	if err = s.orm.Model(&mLocalDelivery.LocalDelivery{}).Where("id = ? AND shop_id = ?", in.Id, s.shopId).
		Update("status", in.Status).Error; err != nil {
		return err
	}
	// 如果是关闭，则不更新下面的内容
	if in.Status == mLocalDelivery.LocalDeliveryStatusClose {
		return err
	}
	// 更新配送区域
	return sDeliveryArea.NewDeliveryArea(s.orm, s.shopId).DeliveryAreaUpdate(in.Areas, in.Id)
}
