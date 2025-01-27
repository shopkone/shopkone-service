package sVariant

import (
	"shopkone-service/internal/module/product/product/mProduct"
)

type ListByProductIdsIn struct {
	ProductIDs []uint
	Keyword    string
	Type       string
}

func (s *sVariant) ListByProductIds(in ListByProductIdsIn) (res []mProduct.Variant, err error) {
	query := s.orm.Model(&mProduct.Variant{}).Where("shop_id = ?", s.shopId)
	query = query.Where("product_id in (?)", in.ProductIDs)
	// 搜索
	if in.Keyword != "" && in.Type != "" {
		if in.Type == "variant_sku" {
			query = query.Where("sku like ?", "%"+in.Keyword+"%")
		} else if in.Type == "name" {
			query = query.Where("name like ?", "%"+in.Keyword+"%")
		}
	}
	query = query.Select("sku", "id", "price", "image_id", "product_id", "name")
	return res, query.Find(&res).Error
}
