package sCollection

import (
	"shopkone-service/internal/module/product/collection/mCollection"
)

func (s *sCollection) CollectionsByProductIds(productIds []uint) (cp []mCollection.CollectionProduct, err error) {
	return cp, s.orm.Model(&cp).
		Where("product_id in ?", productIds).
		Omit("created_at", "deleted_at", "updated_at").
		Find(&cp).Error
}
