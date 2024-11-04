package sTax

import (
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/internal/module/setting/tax/sTax/sCustomerTax"
)

func (s *sTax) TaxRemoveByIds(ids []uint) (err error) {
	// 删除税率
	if err = s.orm.Model(&mTax.Tax{}).
		Where("shop_id = ? AND id IN ?", s.shopId, ids).
		Delete(&mTax.Tax{}).Error; err != nil {
		return err
	}

	// 删除税率区域
	if err = s.TaxZoneRemoveByTaxIds(ids); err != nil {
		return err
	}

	// 删除自定义税率
	return sCustomerTax.NewCustomerTax(s.orm, s.shopId).CustomerTaxRemoveByTaxIds(ids)
}
