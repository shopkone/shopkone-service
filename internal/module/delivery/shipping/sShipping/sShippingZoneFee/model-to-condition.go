package sShippingZoneFee

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

func (s *sShippingZoneFee) ModelToCondition(in mShipping.ShippingZonFeeCondition) (out vo.BaseShippingZoneFeeCondition) {
	out.Min = in.Min
	out.Max = in.Max
	out.Fixed = in.Fixed
	out.First = in.First
	out.FirstFee = in.FirstFee
	out.Next = in.Next
	out.NextFee = in.NextFee
	out.ID = in.ID
	return out
}
