package sTax

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/tax/mTax"
)

func (s *sTax) TaxActive(req vo.TaxActiveReq) (err error) {
	if req.ID == 0 {
		return
	}
	status := mTax.TaxStatusActive
	if !req.Active {
		status = mTax.TaxStatusInactive
	}
	return s.orm.Model(mTax.Tax{}).Where("shop_id = ? AND id = ?", s.shopId, req.ID).
		Update("status", status).Error
}
