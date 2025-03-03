package sOrder

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/order/order/mOrder"
	"shopkone-service/internal/module/order/order/sOrder/sOrderProduct"
	"shopkone-service/internal/module/order/order/sOrder/sOrderShipping"
	"shopkone-service/internal/module/order/order/sOrder/sOrderTax"
	"shopkone-service/internal/module/product/product/sProduct/sProduct"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"
)

// 简易的计算，不涉及优惠券等物品的使用
func (s *sOrder) OrderPreCalPrice(in *vo.OrderCalPreReq) (out vo.OrderCalPreRes, err error) {
	variantService := sOrderProduct.NewOrderVariant(s.orm, s.shopId)

	// 获取商品信息
	variantIds := slice.Map(in.VariantItems, func(index int, item vo.OrderPreBaseVariantItem) uint {
		return item.VariantID
	})
	products, err := variantService.GetProductList(variantIds)
	if err != nil {
		return vo.OrderCalPreRes{}, err
	}
	var variants []sVariant.VariantToOrderOut
	slice.ForEach(products, func(index int, product sProduct.ProductsToOrderOut) {
		slice.ForEach(product.Variants, func(index int, variant sVariant.VariantToOrderOut) {
			variants = append(variants, variant)
		})
	})

	// 计算使用商品优惠券
	discountVariantPrice := slice.Map(variants, func(index int, variant sVariant.VariantToOrderOut) sVariant.VariantToOrderOut {
		inVariant, ok := slice.FindBy(in.VariantItems, func(index int, inVariant vo.OrderPreBaseVariantItem) bool {
			return inVariant.VariantID == variant.ID
		})
		if ok && inVariant.Discount.Price != 0 {
			if inVariant.Discount.Type == mOrder.OrderDiscountTypePercentage {
				variant.Price = variant.Price * (1 - inVariant.Discount.Price/100)
			} else if inVariant.Discount.Type == mOrder.OrderDiscountTypeFixed {
				variant.Price = variant.Price - inVariant.Discount.Price
			}
			if variant.Price < 0 {
				variant.Price = 0
			}
		}
		return variant
	})

	// 计算商品小计
	slice.ForEach(discountVariantPrice, func(index int, variant sVariant.VariantToOrderOut) {
		find, ok := slice.FindBy(in.VariantItems, func(index int, item vo.OrderPreBaseVariantItem) bool {
			return item.VariantID == variant.ID
		})
		if ok {
			out.SumPrice += find.Quantity * variant.Price
		}
	})

	// 计算订单成本价
	slice.ForEach(variants, func(index int, item sVariant.VariantToOrderOut) {
		find, ok := slice.FindBy(in.VariantItems, func(index int, inVariant vo.OrderPreBaseVariantItem) bool {
			return inVariant.VariantID == item.ID
		})
		if ok && item.CostPerItem != nil {
			out.CostPrice += find.Quantity * *item.CostPerItem
		}
	})

	out.Total = out.SumPrice

	// 使用订单优惠券
	if in.Discount.Price != 0 {
		if in.Discount.Type == mOrder.OrderDiscountTypePercentage {
			out.Total = out.Total * (1 - in.Discount.Price/100)
		} else if in.Discount.Type == mOrder.OrderDiscountTypeFixed {
			out.Total = out.Total - in.Discount.Price
		}
	}

	// 获取订单可以使用的运费方案
	if in.CustomerID != 0 && in.Address.Country != "" {
		// 不需要物流则过滤掉
		feeVariants := slice.Filter(discountVariantPrice, func(index int, item sVariant.VariantToOrderOut) bool {
			return item.ShippingRequired
		})
		feesIn := sOrderShipping.FeesByCountryProduct{
			CountryCode:  in.Address.Country,
			ZoneCode:     in.Address.Zone,
			OrderPrice:   out.Total,
			ProductPrice: out.SumPrice,
			Variants: slice.Map(feeVariants, func(index int, v sVariant.VariantToOrderOut) sOrderShipping.FeesProductVariant {
				findIn, ok := slice.FindBy(in.VariantItems, func(index int, item vo.OrderPreBaseVariantItem) bool {
					return item.VariantID == v.ID
				})
				if !ok {
					return sOrderShipping.FeesProductVariant{}
				}
				i := sOrderShipping.FeesProductVariant{}
				i.ID = v.ID
				i.Price = v.Price
				i.ProductId = v.ProductID
				i.Quantity = findIn.Quantity
				i.Weight = v.Weight
				i.WeightUnit = v.WeightUint
				return i
			}),
		}
		list, err := sOrderShipping.NewOrderShipping(s.orm, s.shopId).FeesByCountryProduct(feesIn)
		if err != nil {
			return vo.OrderCalPreRes{}, err
		}
		out.ShippingFeePlans = slice.Map(list.Plans, func(index int, item sOrderShipping.FeesByCountryProductOut) vo.BasePreShippingFeePlan {
			i := vo.BasePreShippingFeePlan{}
			i.ID = item.FeeID
			i.Name = item.FeeName
			i.Price = item.FeePrice
			return i
		})
	}

	// 计算运费
	if in.ShippingFee.Free {
		out.ShippingFee = in.ShippingFee
	} else if in.ShippingFee.Customer {
		out.Total = out.Total + in.ShippingFee.Price
		out.ShippingFee = in.ShippingFee
	}
	if in.ShippingFee.ShippingFeeID != 0 && in.CustomerID != 0 && in.Address.Country != "" && len(out.ShippingFeePlans) > 0 {
		opt, ok := slice.FindBy(out.ShippingFeePlans, func(index int, item vo.BasePreShippingFeePlan) bool {
			return item.ID == in.ShippingFee.ShippingFeeID
		})
		if ok {
			out.ShippingFee = vo.OrderPreBaseShippingFee{
				ShippingFeeID: opt.ID,
				Name:          opt.Name,
				Price:         opt.Price,
			}
		}
	}

	// 计算税
	taxCalIn := sOrderTax.TaxCalIn{
		CountryCode: in.Address.Country,
		ZoneCode:    in.Address.Zone,
		Variants:    discountVariantPrice,
		InVariants:  in.VariantItems,
	}
	taxes, err := sOrderTax.NewOrderTax(s.orm, s.shopId).TaxCal(taxCalIn)
	if err != nil {
		return vo.OrderCalPreRes{}, err
	}
	out.Taxes = slice.Map(taxes, func(index int, item sOrderTax.TaxCalOut) vo.BasePreTaxDetail {
		i := vo.BasePreTaxDetail{}
		i.Rate = item.TaxRate
		i.Price = item.Tax
		i.Name = item.TaxName
		out.Total = out.Total + item.Tax
		return i
	})

	return out, err
}
