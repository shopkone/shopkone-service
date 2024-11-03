package sTax

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sTax) TaxUpdate(in vo.TaxUpdateReq) (err error) {
	var data mTax.Tax
	data.TaxRate = in.TaxRate
	data.HasNote = in.HasNote
	data.Note = in.Note

	// 更新税率
	if err = s.orm.Model(&data).Where("shop_id = ? AND id = ?", s.shopId, in.ID).
		Select("tax_rate", "has_note", "note").Updates(data).Error; err != nil {
		return err
	}

	// 更新区域税率
	if err = s.TaxZoneUpdate(in.Zones, in.ID); err != nil {
		return err
	}

	return err
}
