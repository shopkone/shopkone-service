package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/product/iProduct"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"
)

// 商品是否全都跟踪库存
func (s *sProduct) IsAllTrack(variantIds []uint) (bool, error) {
	// 根据变体id获取商品id
	list, err := sVariant.NewVariant(s.orm, s.shopId).ListByIds(variantIds, false)
	if err != nil {
		return false, err
	}

	if len(list) != len(variantIds) {
		return false, nil
	}

	productIds := slice.Map(list, func(index int, item iProduct.VariantListByIdOut) uint {
		return item.ProductId
	})
	productIds = slice.Unique(productIds)

	// 查找商品是否全都跟踪库存
	var count int64
	query := s.orm.Model(&mProduct.Product{}).Where("shop_id = ? AND id in (?)", s.shopId, productIds)
	query = query.Where("inventory_tracking = ?", true)
	err = query.Count(&count).Error
	return count == int64(len(productIds)), err
}
