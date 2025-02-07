package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/mProduct"
)

func (s *sProduct) UpdateProductOptions(productId uint, ProductOptions []vo.ProductOption) (err error) {
	// 删除旧的options
	if err = s.orm.Where("product_id = ? AND shop_id = ?", productId, s.shopId).
		Unscoped().
		Delete(&mProduct.ProductOption{}).Error; err != nil {
		return err
	}

	// 如果没有options则直接返回
	if len(ProductOptions) == 0 {
		return nil
	}

	// 创建新的 options
	images := slice.Map(ProductOptions, func(index int, item vo.ProductOption) mProduct.ProductOption {
		i := mProduct.ProductOption{}
		i.Values = item.Values
		i.Label = item.Label
		i.ImageId = item.ImageId
		i.ShopId = s.shopId
		i.ProductId = productId
		return i
	})
	return s.orm.Create(&images).Error
}
