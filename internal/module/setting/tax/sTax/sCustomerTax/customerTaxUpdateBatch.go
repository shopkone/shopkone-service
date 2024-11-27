package sCustomerTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/utility/handle"
)

func (s *sCustomerTax) CustomerTaxUpdateBatch(list, oldList []mTax.CustomerTax, baseList []vo.BaseCustomerTax) (err error) {
	// 更新税区域
	if len(list) == 0 {
		return err
	}
	customerIds := slice.Map(list, func(index int, item mTax.CustomerTax) uint {
		return item.ID
	})
	var zones []mTax.CustomerTaxZone
	slice.ForEach(list, func(index int, customer mTax.CustomerTax) {
		base, ok := slice.FindBy(baseList, func(index int, base vo.BaseCustomerTax) bool {
			return customer.ID == base.ID
		})
		if ok {
			slice.ForEach(base.Zones, func(index int, item vo.BaseCustomerTaxZone) {
				z := mTax.CustomerTaxZone{
					CustomerTaxID: customer.ID,
					Name:          item.Name,
					TaxRate:       item.TaxRate,
					CountryCode:   item.CountryCode,
					ZoneCode:      item.ZoneCode,
				}
				z.ID = item.ID
				z.ShopId = s.shopId
				zones = append(zones, z)
			})
		}
	})
	if err = s.CustomerZoneUpdate(zones, customerIds); err != nil {
		return err
	}
	// 更新税
	list = slice.Filter(list, func(index int, customer mTax.CustomerTax) bool {
		old, ok := slice.FindBy(oldList, func(index int, old mTax.CustomerTax) bool {
			return old.ID == customer.ID
		})
		if !ok {
			return false
		}
		return s.CustomerTaxIsChange(old, customer)
	})
	if len(list) == 0 {
		return err
	}
	batchUpdate := handle.BatchUpdateByIdIn{
		Orm:    s.orm,
		ShopID: s.shopId,
		Query:  []string{"collection_id", "type"},
	}
	return handle.BatchUpdateById(batchUpdate, &list)
}
