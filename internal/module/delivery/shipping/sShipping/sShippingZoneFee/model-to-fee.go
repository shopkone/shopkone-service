package sShippingZoneFee

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

func (s *sShippingZoneFee) ModelToFee(in mShipping.ShippingZoneFee) (out vo.BaseShippingZoneFee) {
	out.Name = in.Name
	out.WeightUnit = in.WeightUnit
	out.Type = in.Type
	out.CurrencyCode = in.CurrencyCode
	out.Rule = in.Rule
	out.Remark = in.Remark
	out.ID = in.ID
	out.Cod = in.Cod
	return out
}
