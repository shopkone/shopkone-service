package sCustomerTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sCustomerTax) CustomerTaxRemoveByTaxIds(taxIds []uint) error {
	// 获取自定义税率ids
	var customers []mTax.CustomerTax
	if err := s.orm.Model(&customers).Where("tax_id IN (?) AND shop_id = ?", taxIds, s.shopId).
		Select("id").Find(&customers).Error; err != nil {
		return err
	}

	if len(customers) == 0 {
		return nil
	}

	customerIds := slice.Map(customers, func(index int, item mTax.CustomerTax) uint {
		return item.ID
	})

	// 删除税率id
	if err := s.orm.Where("id IN (?) AND shop_id = ?", customerIds, s.shopId).
		Delete(&mTax.CustomerTax{}).Error; err != nil {
		return err
	}

	// 删除说率区域
	return s.CustomerZoneRemove(customerIds)
}
