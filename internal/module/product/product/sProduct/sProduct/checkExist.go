package sProduct

import (
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/utility/code"
)

func (s *sProduct) CheckExist(ids []uint) (err error) {
	var count int64
	if err = s.orm.Model(&mProduct.Product{}).
		Where("shop_id = ? AND id IN ?", s.shopId, ids).
		Count(&count).Error; err != nil {
		return err
	}
	if count != int64(len(ids)) {
		return code.PartProductNoFound
	}
	return nil
}
