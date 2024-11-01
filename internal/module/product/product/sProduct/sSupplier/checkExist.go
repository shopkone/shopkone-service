package sSupplier

import (
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/utility/code"
)

func (s *sSupplier) CheckExist(id uint) error {
	var count int64
	query := s.orm.Model(&mProduct.Supplier{}).Where("shop_id = ? AND id = ?", s.shopId, id)
	if err := query.Count(&count).Error; err != nil {
		return err
	}
	if count <= 0 {
		return code.ErrSupplierNotFound
	}
	return nil
}
