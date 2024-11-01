package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/utility/handle"
)

func (s *sVariant) UpdateVariants(variants []mProduct.Variant, oldVariants []mProduct.Variant) error {
	// 获取变更了的变体
	variants = slice.Filter(variants, func(index int, item mProduct.Variant) bool {
		find, ok := slice.FindBy(oldVariants, func(index int, i mProduct.Variant) bool {
			return item.ID == i.ID
		})
		if !ok {
			return false
		}
		// 判断是否有变更
		return s.IsChanged(find, item)
	})
	// 更新多条记录
	if len(variants) > 0 {
		variants = slice.Map(variants, func(index int, item mProduct.Variant) mProduct.Variant {
			item.CanCreateId = true
			return item
		})
		batchUpdateIn := handle.BatchUpdateByIdIn{
			Orm:    s.orm,
			Query:  []string{"cost_per_item", "compare_at_price", "price", "sku", "barcode", "weight", "weight_unit", "image_id", "name", "tax_required", "shipping_required"},
			ShopID: s.shopId,
		}
		if err := handle.BatchUpdateById(batchUpdateIn, &variants); err != nil {
			return err
		}
	}
	return nil
}
