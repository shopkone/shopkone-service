package sOrderShipping

import (
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/utility/handle"
)

type FeeFilterIn struct {
	Fee          mShipping.ShippingZoneFee
	Condition    mShipping.ShippingZonFeeCondition
	TotalWeight  float32
	ProductCount int
	OrderPrice   uint32
	ProductPrice uint32
}

// TODO:测试一下shopline如何处理不同名字的方案的
func (s *sOrderShipping) FeeFilter(in FeeFilterIn) bool {
	fee := in.Fee
	condition := in.Condition

	switch fee.Rule {
	case mShipping.ShippingZoneFeeRuleOrderPrice: // 订单总价
		{
			return inRange(condition.Max, condition.Min, float32(in.OrderPrice))
		}
	case mShipping.ShippingZoneFeeRuleProductPrice: // 商品总价
		{
			return inRange(condition.Max, condition.Min, float32(in.ProductPrice))
		}
	case mShipping.ShippingZoneFeeRuleProductCount:
		{
			return inRange(condition.Max, condition.Min, float32(in.ProductCount))
		}
	case mShipping.ShippingZoneFeeRuleOrderWeight:
		{
			maxWeight := handle.ToKg(condition.Max, fee.WeightUnit)
			minWeight := handle.ToKg(condition.Min, fee.WeightUnit)
			return inRange(maxWeight, minWeight, in.TotalWeight)
		}
	}
	return true
}

func inRange(max, min, target float32) bool {
	if target < max && target >= min {
		return true
	}
	if max == 0 {
		return true
	}
	return false
}
