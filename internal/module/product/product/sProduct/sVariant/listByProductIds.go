package sVariant

import "shopkone-service/internal/module/product/product/mProduct"

func (s *sVariant) ListByProductIds(productIds []uint) (res []mProduct.Variant, err error) {
	query := s.orm.Model(&mProduct.Variant{}).Where("shop_id = ?", s.shopId)
	query = query.Where("product_id in (?)", productIds)
	query = query.Select("sku", "id", "price", "image_id", "product_id", "name")
	return res, query.Find(&res).Error
}
