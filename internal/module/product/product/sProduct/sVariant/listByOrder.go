package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/consts"
	"shopkone-service/internal/module/product/product/mProduct"
)

type VariantToOrderOut struct {
	ProductID        uint
	ID               uint
	Price            uint32
	CostPerItem      *uint32
	CompareAtPrice   *uint32
	WeightUint       consts.WeightUnit
	Weight           *float32
	Sku              string
	Barcode          string
	ImageSrc         string
	Name             []mProduct.VariantName
	TaxRequired      bool
	ShippingRequired bool
}

func (s *sVariant) ListByOrder(variantIds []uint) (out []VariantToOrderOut, err error) {
	var list []mProduct.Variant
	if err = s.orm.Model(&list).
		Where("id in (?)", variantIds).
		Find(&list).Error; err != nil {
		return out, err
	}
	getImagesIn := slice.Map(list, func(index int, item mProduct.Variant) GetImagesIn {
		i := GetImagesIn{}
		i.ImageId = item.ImageId
		i.VariantId = item.ID
		return i
	})
	images, err := s.GetImages(getImagesIn)
	if err != nil {
		return nil, err
	}
	out = slice.Map(list, func(index int, item mProduct.Variant) VariantToOrderOut {
		image, ok := slice.FindBy(images, func(index int, image GetImagesOut) bool {
			return item.ID == image.VariantId
		})
		i := VariantToOrderOut{
			ProductID:      item.ProductId,
			ID:             item.ID,
			Price:          item.Price,
			CostPerItem:    item.CostPerItem,
			CompareAtPrice: item.CompareAtPrice,
			WeightUint:     item.WeightUnit,
			Weight:         item.Weight,
			Sku:            item.Sku,
			Barcode:        item.Barcode,
			//ImageSrc:         "",
			Name:             item.Name,
			TaxRequired:      item.TaxRequired,
			ShippingRequired: item.ShippingRequired,
		}
		if ok {
			i.ImageSrc = image.Image
		}
		return i
	})
	return out, err
}
