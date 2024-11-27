package sCustomerTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sCustomerTax) CustomerTaxCreate(in []mTax.CustomerTax, baseList []vo.BaseCustomerTax) (err error) {
	if len(in) == 0 {
		return nil
	}
	// 创建自定义税率
	in = slice.Map(in, func(index int, item mTax.CustomerTax) mTax.CustomerTax {
		item.ID = 0
		item.ShopId = s.shopId
		return item
	})
	if err = s.orm.Create(in).Error; err != nil {
		return err
	}
	// 创建自定义税率区域
	var zones []mTax.CustomerTaxZone
	slice.ForEach(in, func(index int, customer mTax.CustomerTax) {
		base, ok := slice.FindBy(baseList, func(index int, base vo.BaseCustomerTax) bool {
			if customer.Type == mTax.CustomerTaxTypeCollection {
				return base.CollectionID == customer.CollectionID
			}
			return true
		})
		if !ok {
			return
		}
		slice.ForEach(base.Zones, func(index int, zone vo.BaseCustomerTaxZone) {
			i := mTax.CustomerTaxZone{
				CustomerTaxID: customer.ID,
				Name:          zone.Name,
				TaxRate:       zone.TaxRate,
				CountryCode:   zone.CountryCode,
				ZoneCode:      zone.ZoneCode,
			}
			i.ShopId = s.shopId
			zones = append(zones, i)
		})
	})
	return s.CustomerZoneCreate(zones)
}
