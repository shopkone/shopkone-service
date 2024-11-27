package sCustomerTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/tax/mTax"
)

type CustomerTaxByCountryProductIn struct {
	TaxID         uint
	CountryCode   string
	ZoneCode      string
	CollectionIds []uint
}

type TaxItem struct {
	Name string
	Rate float64
}

type CustomerTaxByCountryProductOut struct {
	TaxName      string
	TaxRate      float64
	CollectionID uint
}

func (s *sCustomerTax) CustomerTaxByCountryProduct(in CustomerTaxByCountryProductIn) (res []CustomerTaxByCountryProductOut, feeTax TaxItem, err error) {
	// 获取区域自定义税率
	var customerTaxAreas []mTax.CustomerTax
	if len(in.CollectionIds) > 0 {
		if err = s.orm.Model(&customerTaxAreas).
			Where("tax_id = ?", in.TaxID).
			Where("collection_id IN ?", in.CollectionIds).
			Where("type = ?", mTax.CustomerTaxTypeCollection).
			Omit("created_at", "updated_at", "deleted_at").
			Find(&customerTaxAreas).Error; err != nil {
			return res, feeTax, err
		}
		customerAreaIds := slice.Map(customerTaxAreas, func(index int, item mTax.CustomerTax) uint {
			return item.ID
		})
		// 获取区域自定义税率详细内容
		var customerTaxAreaZones []mTax.CustomerTaxZone
		if err = s.orm.Model(&customerTaxAreaZones).
			Where("customer_tax_id IN ?", customerAreaIds).
			Where("area_code = ? OR area_code = ?", in.CountryCode, in.ZoneCode).
			Omit("created_at", "updated_at", "deleted_at").
			Find(&customerTaxAreaZones).Error; err != nil {
			return res, feeTax, err
		}
		res = slice.Map(in.CollectionIds, func(index int, collectionId uint) CustomerTaxByCountryProductOut {
			// 找到该集合下的自定义税
			c, ok := slice.FindBy(customerTaxAreas, func(index int, item mTax.CustomerTax) bool {
				return item.CollectionID == collectionId
			})
			if !ok {
				return CustomerTaxByCountryProductOut{}
			}
			// 找到该集合下的税率
			zone, ok := slice.FindBy(customerTaxAreaZones, func(index int, item mTax.CustomerTaxZone) bool {
				return c.ID == item.CustomerTaxID
			})
			if !ok {
				return CustomerTaxByCountryProductOut{}
			}
			i := CustomerTaxByCountryProductOut{}
			i.TaxName = zone.Name
			i.TaxRate = zone.TaxRate
			i.CollectionID = collectionId
			return i
		})
		res = slice.Filter(res, func(index int, item CustomerTaxByCountryProductOut) bool {
			return item.CollectionID > 0
		})
	}

	// 获取运费自定义税率
	var shippingTaxArea mTax.CustomerTax
	if err = s.orm.Model(&shippingTaxArea).
		Where("tax_id = ? AND type = ?", in.TaxID, mTax.CustomerTaxTypeDelivery).
		Omit("created_at", "updated_at", "deleted_at").
		Find(&shippingTaxArea).Error; err != nil {
		return res, feeTax, err
	}
	// 获取运费自定义税率详细内容
	var shippingTaxAreaZone mTax.CustomerTaxZone
	if err = s.orm.Model(&shippingTaxAreaZone).
		Where("customer_tax_id = ?", shippingTaxArea.ID).
		Omit("created_at", "updated_at", "deleted_at").
		Find(&shippingTaxAreaZone).Error; err != nil {
		return res, feeTax, err
	}
	if shippingTaxAreaZone.ID != 0 {
		feeTax = TaxItem{
			Name: shippingTaxAreaZone.Name,
			Rate: shippingTaxAreaZone.TaxRate,
		}
	}

	return res, feeTax, err
}
