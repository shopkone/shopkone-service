package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/product/mProduct"
)

func (s *sProduct) IdsByProduct(productId uint) (ids []uint, err error) {
	var variants []mProduct.Variant
	if err = s.orm.Model(&variants).Where("product_id = ?", productId).
		Select("id").Find(&variants).Error; err != nil {
		return nil, err
	}
	ids = slice.Map(variants, func(index int, item mProduct.Variant) uint {
		return item.ID
	})
	return ids, err
}
