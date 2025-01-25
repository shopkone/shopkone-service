package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/collection/mCollection"
)

func (s *sProduct) GetCollections(productId uint) (collectionIds []uint, err error) {
	var collections []mCollection.CollectionProduct
	if err = s.orm.Model(&collections).Where("product_id = ?", productId).
		Select("collection_id").Find(&collections).Error; err != nil {
		return collectionIds, err
	}
	return slice.Map(collections, func(index int, item mCollection.CollectionProduct) uint {
		return item.CollectionId
	}), nil
}
