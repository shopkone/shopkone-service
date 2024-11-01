package sProduct

import (
	"shopkone-service/internal/module/base/orm/mOrm"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/utility/code"
)

func (s *sProduct) UpdateNameHandler(variants []mProduct.Variant, productId uint) error {
	if len(variants) == 0 {
		return code.NoEmptyVariants
	}
	// 先删除
	query := s.orm.Model(&mProduct.VariantNameHandler{}).Where("product_id=?", productId)
	query = query.Where("shop_id = ?", s.shopId)
	if err := query.Unscoped().Delete(&mProduct.VariantNameHandler{}).Error; err != nil {
		return err
	}
	// 再添加
	var handlers []mProduct.VariantNameHandler
	if len(handlers) == 0 {
		return nil
	}
	for _, variant := range variants {
		for _, name := range variant.Name {
			handlers = append(handlers, mProduct.VariantNameHandler{
				ProductId: variant.ProductId,
				Label:     name.Label,
				Value:     name.Value,
				VariantId: variant.ID,
				Model:     mOrm.Model{ShopId: s.shopId},
			})
		}
	}
	return s.orm.Create(&handlers).Error
}
