package sCollection

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/collection/mCollection"
	"shopkone-service/utility/code"
)

func (s *sCollection) CheckAllExist(collectionIds []uint) (err error) {
	collectionIds = slice.Unique(collectionIds)
	collectionIds = slice.Filter(collectionIds, func(index int, item uint) bool {
		return item > 0
	})

	if len(collectionIds) == 0 {
		return err
	}

	var count int64
	if err = s.orm.Model(&mCollection.ProductCollection{}).Where("shop_id =?", s.shopId).
		Where("id IN (?)", collectionIds).Count(&count).Error; err != nil {
		return err
	}

	if count != int64(len(collectionIds)) {
		return code.ErrCollectionNotFound
	}

	return err
}
