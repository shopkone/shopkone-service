package sOrderShipping

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/consts"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShipping"
	"shopkone-service/internal/module/delivery/shipping/sShipping/sShippingZoneFee"
	"shopkone-service/utility/handle"
)

type FeesProductVariant struct {
	ID         uint
	Price      float32
	Quantity   uint
	Weight     *float32
	WeightUnit consts.WeightUnit
	ProductId  uint
}

type FeesByCountryProduct struct {
	CountryCode  string
	ZoneCode     string
	Variants     []FeesProductVariant
	OrderPrice   float32
	ProductPrice float32
}

type FeesByCountryProductOut struct {
	FeeName  string
	FeeID    uint
	FeePrice float32
}

func (s *sOrderShipping) FeesByCountryProduct(in FeesByCountryProduct) (out []FeesByCountryProductOut, err error) {
	productIds := slice.Map(in.Variants, func(index int, item FeesProductVariant) uint {
		return item.ProductId
	})
	productIds = slice.Unique(productIds)

	// 获取费用
	feesIn := sShipping.ShippingFeesByCountryProductIn{
		CountryCode: in.CountryCode,
		ProductIDs:  productIds,
		ZoneCode:    in.ZoneCode,
	}
	fees, err := sShipping.NewShipping(s.orm, s.shopId).FeesByCountryProduct(feesIn)
	if err != nil {
		return out, err
	}

	// 获取费用条件
	feeIds := slice.Map(fees, func(index int, item mShipping.ShippingZoneFee) uint {
		return item.ID
	})
	feeIds = slice.Unique(feeIds)
	feeConditions, err := sShippingZoneFee.NewShippingZoneFee(s.orm, s.shopId).ConditionsByFeeIds(feeIds)
	if err != nil {
		return nil, err
	}

	// 商品总数
	var productCount int
	slice.ForEach(in.Variants, func(index int, item FeesProductVariant) {
		productCount = productCount + int(item.Quantity)
	})
	// 商品总重
	var totalWeight float32
	slice.ForEach(in.Variants, func(index int, item FeesProductVariant) {
		if item.Weight == nil {
			return
		}
		totalWeight = totalWeight + handle.ToKg(*item.Weight, item.WeightUnit)
	})

	slice.ForEach(feeConditions, func(index int, condition mShipping.ShippingZonFeeCondition) {
		fee, ok := slice.FindBy(fees, func(index int, fee mShipping.ShippingZoneFee) bool {
			return fee.ID == condition.ShippingZoneFeeId
		})
		if !ok {
			return
		}
		feeIn := FeeCalIn{
			Fee:          fee,
			Condition:    condition,
			TotalWeight:  totalWeight,
			ProductCount: productCount,
			OrderPrice:   in.OrderPrice,
			ProductPrice: in.ProductPrice,
		}
		if s.FeeFilter(feeIn) {
			i := FeesByCountryProductOut{}
			i.FeeID = fee.ID
			i.FeeName = fee.Name
			i.FeePrice = s.FeeCal(feeIn)
			out = append(out, i)
		}
	})

	return out, err
}
