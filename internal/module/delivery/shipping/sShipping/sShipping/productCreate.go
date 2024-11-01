package sShipping

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

func (s *sShipping) ProductCreate(shippingId uint, productIDs []uint) error {
	if len(productIDs) == 0 {
		return nil
	}

	// 删除旧的商品
	if err := s.ProductRemove(productIDs); err != nil {
		return err
	}

	// 创建适用商品
	shippingProduct := slice.Map(productIDs, func(index int, item uint) mShipping.ShippingProduct {
		i := mShipping.ShippingProduct{}
		i.ShippingId = shippingId
		i.ProductId = item
		i.ShopId = s.shopId
		return i
	})
	return s.orm.Create(&shippingProduct).Error
}
