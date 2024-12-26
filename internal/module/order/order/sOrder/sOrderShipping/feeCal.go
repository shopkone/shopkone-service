package sOrderShipping

import (
	"math"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/utility/handle"
)

type FeeCalInItem struct {
	Fee          mShipping.ShippingZoneFee
	Condition    mShipping.ShippingZonFeeCondition
	TotalWeight  float32
	ProductCount int
}

func (s *sOrderShipping) FeeCalItem(in FeeCalInItem) (price uint32) {
	fee := in.Fee
	condition := in.Condition
	productCount := in.ProductCount

	switch fee.Type {
	case mShipping.ShippingZoneFeeTypeFixed:
		{ // 固定运费
			return condition.Fixed
		}
	case mShipping.ShippingZoneFeeTypeWeight:
		{ // 按重量计费
			firstWeight := handle.ToKg(condition.First, fee.WeightUnit)
			nextWeight := handle.ToKg(condition.Next, fee.WeightUnit)
			firstPrice := condition.FirstFee
			var remainPrice uint32
			remain := in.TotalWeight - firstWeight
			if remain > 0 {
				r := float32(math.Ceil(float64(remain / nextWeight)))
				remainPrice = condition.NextFee * uint32(r)
			}
			return firstPrice + remainPrice
		}
	case mShipping.ShippingZoneFeeTypeCount:
		{ // 按商品数量计费
			firstPrice := condition.FirstFee
			var remainPrice uint32
			remain := productCount - int(condition.First)
			if remain > 0 {
				r := float32(math.Ceil(float64(remain / int(condition.Next))))
				remainPrice = condition.NextFee * uint32(r)
			}
			return firstPrice + remainPrice
		}
	}

	return 0
}
