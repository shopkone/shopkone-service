package sTax

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sTax) TaxZonesByTaxId(taxId uint) (out []vo.BaseTaxZone, err error) {
	var list []mTax.TaxZone
	if err = s.orm.Model(&list).Where("tax_id = ?", taxId).Where("shop_id = ?", s.shopId).
		Omit("created_at", "deleted_at", "updated_at").Find(&list).Error; err != nil {
		return
	}
	out = slice.Map(list, func(index int, item mTax.TaxZone) vo.BaseTaxZone {
		return vo.BaseTaxZone{
			ID:       item.ID,
			Name:     item.Name,
			ZoneCode: item.ZoneCode,
			TaxRate:  item.TaxRate,
		}
	})
	return out, err
}
