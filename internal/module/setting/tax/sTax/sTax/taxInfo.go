package sTax

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/tax/mTax"
	"shopkone-service/internal/module/setting/tax/sTax/sCustomerTax"
)

func (s *sTax) TaxInfo(id uint) (res vo.TaxInfoRes, err error) {
	var data mTax.Tax
	if err = s.orm.Model(&data).Where("id = ?", id).
		Omit("created_at", "updated_at", "deleted_at", "shop_id").
		First(&data).Error; err != nil {
		return res, err
	}

	res.ID = data.ID
	res.TaxRate = data.TaxRate
	res.CountryCode = data.CountryCode
	res.HasNote = data.HasNote
	res.Note = data.Note
	res.Status = data.Status

	if res.Zones, err = s.TaxZonesByTaxId(res.ID); err != nil {
		return vo.TaxInfoRes{}, err
	}

	if res.Customers, err = sCustomerTax.NewCustomerTax(s.orm, s.shopId).CustomerTaxListByTaxId(res.ID); err != nil {
		return vo.TaxInfoRes{}, err
	}

	return res, err
}
