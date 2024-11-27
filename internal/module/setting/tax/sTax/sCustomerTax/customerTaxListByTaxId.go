package sCustomerTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sCustomerTax) CustomerTaxListByTaxId(taxId uint) (res []vo.BaseCustomerTax, err error) {
	// 查询自定义税列表
	var customerTaxList []mTax.CustomerTax
	if err = s.orm.Model(&customerTaxList).
		Where("tax_id = ? AND shop_id = ?", taxId, s.shopId).
		Omit("created_at", "updated_at", "deleted_at", "shop_id").Find(&customerTaxList).Error; err != nil {
		return nil, err
	}
	// 查询自定义税区域
	customerTaxIDs := slice.Map(customerTaxList, func(index int, item mTax.CustomerTax) uint {
		return item.ID
	})
	var zones []mTax.CustomerTaxZone
	if err = s.orm.Model(&zones).
		Where("customer_tax_id IN ?", customerTaxIDs).
		Omit("created_at", "updated_at", "deleted_at", "shop_id").Find(&zones).Error; err != nil {
		return nil, err
	}
	res = slice.Map(customerTaxList, func(index int, item mTax.CustomerTax) vo.BaseCustomerTax {
		i := vo.BaseCustomerTax{}
		i.ID = item.ID
		i.Type = item.Type
		i.CollectionID = item.CollectionID
		currentZones := slice.Filter(zones, func(index int, zone mTax.CustomerTaxZone) bool {
			return zone.CustomerTaxID == item.ID
		})
		i.Zones = slice.Map(currentZones, func(index int, zone mTax.CustomerTaxZone) vo.BaseCustomerTaxZone {
			return vo.BaseCustomerTaxZone{
				ID:          zone.ID,
				Name:        zone.Name,
				ZoneCode:    zone.ZoneCode,
				CountryCode: zone.CountryCode,
				TaxRate:     zone.TaxRate,
			}
		})
		return i
	})
	return res, err
}
