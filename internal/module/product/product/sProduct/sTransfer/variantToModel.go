package sTransfer

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/mProduct"
)

func (s *sProductTransfer) VariantToModel(in vo.BaseVariant, productId uint, needId bool) (out mProduct.Variant) {
	out = mProduct.Variant{
		Name:             in.Name,
		Price:            in.Price,
		Sku:              in.Sku,
		Barcode:          in.Barcode,
		ImageId:          in.ImageId,
		Weight:           in.Weight,
		WeightUnit:       in.WeightUnit,
		CompareAtPrice:   in.CompareAtPrice,
		CostPerItem:      in.CostPerItem,
		ProductId:        productId,
		TaxRequired:      in.TaxRequired,
		ShippingRequired: in.ShippingRequired,
	}
	if needId {
		out.ID = in.Id
	}
	if productId == 0 {
		return out
	}
	return out
}
