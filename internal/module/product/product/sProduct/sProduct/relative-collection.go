package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/collection/mCollection"
)

func (s *sProduct) RelativeCollection(collectionIds []uint, productId uint) (err error) {
	// 删除之前的关联
	if err = s.orm.Model(mCollection.CollectionProduct{}).
		Where("product_id = ?", productId).
		Unscoped().Delete(&mCollection.CollectionProduct{}).Error; err != nil {
		return err
	}
	// 手动关联
	if len(collectionIds) > 0 {
		data := slice.Map(collectionIds, func(index int, item uint) mCollection.CollectionProduct {
			i := mCollection.CollectionProduct{}
			i.ProductId = productId
			i.CollectionId = item
			i.ShopId = s.shopId
			return i
		})
		if err = s.orm.Create(&data).Error; err != nil {
			return err
		}
	}
	// 自动关联
	return err
}
