package sShipping

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

func (s *sShipping) ProductUpdate(shippingId uint, newProductIDs []uint) error {
	// 获取旧的商品
	var oldProducts []mShipping.ShippingProduct
	if err := s.orm.Where("shipping_id = ? AND shop_id = ?", shippingId, s.shopId).
		Select("product_id").Find(&oldProducts).Error; err != nil {
		return err
	}
	oldProductIds := slice.Map(oldProducts, func(index int, item mShipping.ShippingProduct) uint {
		return item.ProductId
	})

	// 找出要删除的商品
	deleteProductIds := slice.Filter(oldProductIds, func(index int, item uint) bool {
		return slice.Contain(newProductIDs, item) == false
	})
	if err := s.ProductRemove(deleteProductIds); err != nil {
		return err
	}

	// 找出要添加的商品
	addProductIds := slice.Filter(newProductIDs, func(index int, item uint) bool {
		return slice.Contain(oldProductIds, item) == false
	})
	return s.ProductCreate(shippingId, addProductIds)
}
