package sTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/base/resource"
	"shopkone-service/internal/module/base/resource/mResource"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sTax) TaxCreate(countryCodes []string) (err error) {
	countryCodes = slice.Unique(countryCodes)

	// 获取已经存在的countryCodes
	var taxs []mTax.Tax
	if err = s.orm.Model(&taxs).Where("shop_id = ? AND country_code IN ?", s.shopId, countryCodes).
		Select("country_code").Find(&taxs).Error; err != nil {
		return err
	}

	// 过滤掉已经存在的
	countryCodes = slice.Filter(countryCodes, func(index int, code string) bool {
		_, ok := slice.FindBy(taxs, func(index int, tax mTax.Tax) bool {
			return tax.CountryCode == code
		})
		return !ok
	})

	// 创建
	if len(countryCodes) == 0 {
		return err
	}

	taxs = slice.Map(countryCodes, func(index int, code string) mTax.Tax {
		find, ok := slice.FindBy(resource.Taxs, func(index int, item mResource.Tax) bool {
			return item.CountryCode == code
		})
		i := mTax.Tax{}
		i.CountryCode = code
		i.ShopId = s.shopId
		if ok {
			i.TaxRate = find.Tax
		}
		return i
	})
	return s.orm.Create(&taxs).Error
}
