package sTransfer

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/mProduct"
)

func (s *sProductTransfer) ModelToVariant(in mProduct.Variant) (out vo.BaseVariant) {
	out = vo.BaseVariant{
		Name:             in.Name,
		Price:            in.Price,
		Sku:              in.Sku,
		Barcode:          in.Barcode,
		ImageId:          in.ImageId,
		Weight:           in.Weight,
		WeightUnit:       in.WeightUnit,
		CompareAtPrice:   in.CompareAtPrice,
		CostPerItem:      in.CostPerItem,
		TaxRequired:      in.TaxRequired,
		ShippingRequired: in.ShippingRequired,
	}
	out.Id = in.ID
	return out
}
