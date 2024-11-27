package sTax

import (
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/internal/module/setting/tax/sTax/sCustomerTax"
)

type TaxZoneByAreaCodeIn struct {
	ZoneCode string
	TaxID    uint
}

func (s *sTax) TaxZoneByAreaCode(in TaxZoneByAreaCodeIn) (out sCustomerTax.TaxItem, err error) {
	if in.ZoneCode == "" {
		return out, err
	}
	var data mTax.TaxZone
	if err = s.orm.Where("tax_id = ? AND zone_code = ?", in.TaxID, in.ZoneCode).
		Omit("created_at", "updated_at", "deleted_at").
		Find(&data).Error; err != nil {
		return out, err
	}
	out = sCustomerTax.TaxItem{
		Name: data.Name,
		Rate: data.TaxRate,
	}
	return out, err
}
