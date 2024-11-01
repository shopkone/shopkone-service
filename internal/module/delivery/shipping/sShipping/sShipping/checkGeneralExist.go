package sShipping

import (
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/utility/code"
)

func (s *sShipping) CheckGeneralExist() (err error) {
	// 判断通用方案是否已经存在
	var count int64
	if err = s.orm.Model(&mShipping.Shipping{}).
		Where("type = ? AND shop_id = ?", mShipping.GeneralExpressDelivery, s.shopId).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return code.GeneralShippingExist
	}
	return err
}
