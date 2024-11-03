package sCustomerTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sCustomerTax) CustomerTaxRemove(list []mTax.CustomerTax) (err error) {
	ids := slice.Map(list, func(index int, item mTax.CustomerTax) uint {
		return item.ID
	})
	if len(ids) == 0 {
		return err
	}
	if err = s.orm.Where("shop_id = ? AND id IN ?", s.shopId, ids).
		Delete(&mTax.CustomerTax{}).Error; err != nil {
		return err
	}
	return s.CustomerZoneRemove(ids)
}
