package sOrderShipping

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/delivery/shipping/mShipping"
	"shopkone-service/utility/handle"
)

type FeeCalsPerProductIn struct {
	Fee          mShipping.ShippingZoneFee
	Condition    mShipping.ShippingZonFeeCondition
	Variant      FeesProductVariant
	IsGeneralFee bool
}

type FeeCalsPerProductIn1 struct {
	Fee       mShipping.ShippingZoneFee
	Condition mShipping.ShippingZonFeeCondition
	Variants  []FeesProductVariant
	Price     uint32
}

type FeeCalsPerProductOut struct {
	Plans  []FeesByCountryProductOut
	Groups []FeeCalsPerProductIn1 // 如果时组合几个运费方案一起时才会展现
}

// 计算运费方案
func (s *sOrderShipping) FeeCalsPerProduct(in []FeeCalsPerProductIn, allVariant []FeesProductVariant) (out FeeCalsPerProductOut) {
	// 首先过滤掉已经有的
	in = slice.Filter(in, func(index int, item FeeCalsPerProductIn) bool {
		if !item.IsGeneralFee {
			return true
		}
		_, has := slice.FindBy(in, func(index int, i FeeCalsPerProductIn) bool {
			return i.Variant.ID == item.Variant.ID && !i.IsGeneralFee
		})
		return !has
	})
	// 合并同一个运费方案
	var list []FeeCalsPerProductIn1
	slice.ForEach(in, func(index int, item FeeCalsPerProductIn) {
		find, has := slice.FindBy(list, func(index int, i FeeCalsPerProductIn1) bool {
			return i.Fee.ID == item.Fee.ID
		})
		if !has {
			list = append(list, FeeCalsPerProductIn1{
				Fee:       item.Fee,
				Condition: item.Condition,
				Variants:  []FeesProductVariant{item.Variant},
			})
		} else {
			find.Variants = append(find.Variants, item.Variant)
			list = slice.Map(list, func(index int, i FeeCalsPerProductIn1) FeeCalsPerProductIn1 {
				if i.Fee.ID == find.Fee.ID {
					return find
				}
				return i
			})
		}
	})
	list = slice.Map(list, func(index int, item FeeCalsPerProductIn1) FeeCalsPerProductIn1 {
		var weight float32
		slice.ForEach(item.Variants, func(index int, item FeesProductVariant) {
			if item.Weight == nil {
				return
			}
			w := handle.ToKg(*item.Weight, item.WeightUnit) * float32(item.Quantity)
			weight = weight + w
		})
		falCalIn := FeeCalInItem{
			Condition:    item.Condition,
			Fee:          item.Fee,
			ProductCount: len(item.Variants),
			TotalWeight:  weight,
		}
		item.Price = item.Price + s.FeeCalItem(falCalIn)
		g.Dump(item.Price, "PRICE", item.Fee.Name)
		return item
	})
	canBackPlans := slice.Filter(list, func(index int, item FeeCalsPerProductIn1) bool {
		return slice.Every(allVariant, func(index int, v FeesProductVariant) bool {
			_, ok := slice.FindBy(item.Variants, func(index int, av FeesProductVariant) bool {
				return av.ID == v.ID
			})
			return ok
		})
	})
	if len(canBackPlans) > 0 {
		out.Plans = slice.Map(canBackPlans, func(index int, item FeeCalsPerProductIn1) FeesByCountryProductOut {
			return FeesByCountryProductOut{
				FeeName:  item.Fee.Name,
				FeeID:    item.Fee.ID,
				FeePrice: item.Price,
			}
		})
		return out
	} else {
		i := FeesByCountryProductOut{
			FeeName:  "",
			FeeID:    0,
			FeePrice: 0,
		}
		slice.ForEach(allVariant, func(idx int, item FeesProductVariant) {
			plans := slice.Filter(list, func(index int, f FeeCalsPerProductIn1) bool {
				_, ok := slice.FindBy(f.Variants, func(index int, i FeesProductVariant) bool {
					return i.ID == item.ID
				})
				if ok {
					return true
				}
				return false
			})
			var minPlan FeeCalsPerProductIn1
			// 从所有计划找到最便宜的计划
			slice.ForEach(plans, func(index int, item FeeCalsPerProductIn1) {
				if index == 0 {
					minPlan = item
				} else if minPlan.Price > item.Price {
					minPlan = item
				}
				out.Groups = append(out.Groups, minPlan)
			})
			i.FeePrice = i.FeePrice + minPlan.Price
			if idx == 0 {
				i.FeeName = minPlan.Fee.Name
			} else if minPlan.Fee.Name == i.FeeName {
				i.FeeName = minPlan.Fee.Name
			} else {
				i.FeeName = "Shipping"
			}
			i.FeeID = 1 // 随便取的一个值
		})
		out.Plans = []FeesByCountryProductOut{i}
		return out
	}
}
