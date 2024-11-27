package sTax

import (
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/internal/module/setting/tax/sTax/sCustomerTax"
)

type TaxByCountryProductIn struct {
	CountryCode   string
	ZoneCode      string
	CollectionIDs []uint
}

type TaxByCountryProductOut struct {
	Base             sCustomerTax.TaxItem
	Zone             sCustomerTax.TaxItem
	CustomerShipping sCustomerTax.TaxItem
	Customer         []sCustomerTax.CustomerTaxByCountryProductOut
}

func (s *sTax) TaxByCountryProduct(in TaxByCountryProductIn) (out TaxByCountryProductOut, err error) {
	var tax mTax.Tax
	if err = s.orm.Where("country_code = ?", in.CountryCode).
		Omit("created_at", "updated_at", "deleted_at").
		Find(&tax).Error; err != nil {
		return out, err
	}
	if tax.Status == mTax.TaxStatusInactive {
		return out, nil
	}
	out.Base.Name = ""
	out.Base.Rate = tax.TaxRate

	// 查找自定义税率
	customerIn := sCustomerTax.CustomerTaxByCountryProductIn{
		CollectionIds: in.CollectionIDs,
		CountryCode:   in.CountryCode,
		TaxID:         tax.ID,
		ZoneCode:      in.ZoneCode,
	}
	out.Customer, out.CustomerShipping, err = sCustomerTax.NewCustomerTax(s.orm, s.shopId).CustomerTaxByCountryProduct(customerIn)
	if err != nil {
		return out, err
	}

	// 查找区域税率
	taxZoneIn := TaxZoneByAreaCodeIn{
		TaxID:    tax.ID,
		ZoneCode: in.ZoneCode,
	}
	out.Zone, err = s.TaxZoneByAreaCode(taxZoneIn)
	if err != nil {
		return out, err
	}

	return out, err
}
