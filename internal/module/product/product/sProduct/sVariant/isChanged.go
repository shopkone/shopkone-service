package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/product/mProduct"
)

func (s *sVariant) IsChanged(newVariant, oldVariant mProduct.Variant) bool {
	isChangeCostPerItem := newVariant.CostPerItem != oldVariant.CostPerItem
	isChangeCompareAtPrice := newVariant.CompareAtPrice != oldVariant.CompareAtPrice
	isChangePrice := newVariant.Price != oldVariant.Price
	isChangeSku := newVariant.Sku != oldVariant.Sku
	isChangeBarcode := newVariant.Barcode != oldVariant.Barcode
	isChangeWeight := newVariant.Weight != oldVariant.Weight
	isChangeWeightUint := newVariant.WeightUnit != oldVariant.WeightUnit
	isChangeImage := newVariant.ImageId != oldVariant.ImageId
	isChangeTax := newVariant.TaxRequired != oldVariant.TaxRequired
	isChangeShipping := newVariant.ShippingRequired != oldVariant.ShippingRequired
	isChangeName := slice.Some(newVariant.Name, func(index int, item mProduct.VariantName) bool {
		_, has := slice.FindBy(oldVariant.Name, func(index int, n mProduct.VariantName) bool {
			return item.Label == n.Label && item.Value == n.Value
		})
		return !has
	})
	return isChangeCostPerItem || isChangeCompareAtPrice || isChangeSku || isChangeBarcode || isChangeWeight || isChangeWeightUint || isChangePrice || isChangeName || isChangeImage || isChangeTax || isChangeShipping
}
