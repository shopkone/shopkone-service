package sShipping

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZone"
)

func (s *sShipping) ShippingUpdate(in vo.ShippingUpdateReq) (err error) {
	if err = s.Check(in.BaseShipping, in.ID); err != nil {
		return err
	}

	if in.Type == mShipping.CustomerExpressDelivery {
		// 更新物流方案
		if err = s.orm.Model(&mShipping.Shipping{}).Where("shop_id = ? AND id = ?", s.shopId, in.ID).
			Update("name", in.Name).Error; err != nil {
			return err
		}

		// 更新商品ids
		if err = s.ProductUpdate(in.ID, in.ProductIDs); err != nil {
			return err
		}

		// 更新发货地点
		if err = s.LocationUpdate(in.Id, in.LocationIDs); err != nil {
			return err
		}
	}

	// 更新区域数据
	if err = sShippingZone.NewShippingZone(s.shopId, s.orm).ZoneUpdate(in.Zones, in.ID); err != nil {
		return err
	}

	return err
}
