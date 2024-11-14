package sShipping

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZone"
)

func (s *sShipping) ShippingCreate(in vo.ShippingCreateReq) (id uint, err error) {
	if err = s.Check(in.BaseShipping, 0); err != nil {
		return 0, err
	}

	// 创建物流方案
	shipping := mShipping.Shipping{}
	shipping.ShopId = s.shopId
	shipping.Name = in.Name
	if shipping.Type == mShipping.GeneralExpressDelivery {
		shipping.Name = ""
	}
	shipping.Type = in.Type
	if err = s.orm.Create(&shipping).Error; err != nil {
		return 0, err
	}

	if in.Type == mShipping.CustomerExpressDelivery {
		// 创建商品ids
		if err = s.ProductCreate(shipping.ID, in.ProductIDs); err != nil {
			return 0, err
		}
		// 创建发货地点ids
		if err = s.LocationCreate(shipping.ID, in.LocationIDs); err != nil {
			return 0, err
		}
	}

	// 创建收货区域
	if err = sShippingZone.NewShippingZone(s.shopId, s.orm).ZoneCreate(in.Zones, shipping.ID); err != nil {
		return 0, err
	}

	return shipping.ID, nil
}
