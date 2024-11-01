package sTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sTax) TaxList() (res []vo.TaxListRes, err error) {
	var list []mTax.Tax
	if err = s.orm.Model(&list).Where("shop_id = ?", s.shopId).
		Select("id", "tax_rate", "country_code", "status").Find(&list).Error; err != nil {
		return res, err
	}

	res = slice.Map(list, func(index int, tax mTax.Tax) vo.TaxListRes {
		return vo.TaxListRes{
			Id:          tax.ID,
			CountryCode: tax.CountryCode,
			Status:      tax.Status,
			TaxRate:     tax.TaxRate,
		}
	})
	return res, nil
}
