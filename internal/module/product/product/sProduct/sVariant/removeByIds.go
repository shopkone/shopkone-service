package sVariant

import (
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/product/purchase/mPurchase"
)

func (s *sVariant) RemoveByIds(ids []uint) error {
	// 删除变体
	err := s.orm.Where("id IN ? AND shop_id = ?", ids, s.shopId).
		Delete(&mProduct.Variant{}).Error
	if err != nil {
		return err
	}
	// 禁用采购单项，直接在这里操作是因为循环引用
	query := s.orm.Model(&mPurchase.PurchaseItem{})
	query = query.Where("shop_id = ? AND variant_id IN ?", s.shopId, ids)
	query = query.Update("active", false)
	if query.Error != nil {
		return err
	}
	// 删除库存
	return sInventory.NewInventory(s.orm, s.shopId).RemoveByVariantIds(ids, 0)
}
