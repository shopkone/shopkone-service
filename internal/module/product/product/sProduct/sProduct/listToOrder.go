package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"
)

type ProductsToOrderOut struct {
	ID                uint
	Title             string
	Spu               string
	Vendor            string
	Tags              []string
	Type              mProduct.VariantType
	Category          string
	FirstImageSrc     string
	Variants          []sVariant.VariantToOrderOut
	InventoryTracking bool
}

func (s *sProduct) ListToOrder(variantIds []uint) ([]ProductsToOrderOut, error) {
	// 获取变体列表
	variants, err := sVariant.NewVariant(s.orm, s.shopId).ListByOrder(variantIds)
	if err != nil {
		return nil, err
	}
	productIds := slice.Map(variants, func(index int, item sVariant.VariantToOrderOut) uint {
		return item.ProductID
	})

	// 获取商品列表
	var list []mProduct.Product
	if err = s.orm.Model(&list).Where("id in (?)", productIds).
		Omit("created_at", "updated_at", "deleted_at").
		Find(&list).Error; err != nil {
		return nil, err
	}

	// 组装数据
	out := slice.Map(list, func(index int, item mProduct.Product) ProductsToOrderOut {
		i := ProductsToOrderOut{
			ID:                item.ID,
			Title:             item.Title,
			Spu:               item.Spu,
			Vendor:            item.Vendor,
			Tags:              item.Tags,
			Type:              item.VariantType,
			InventoryTracking: item.InventoryTracking,
			//Category:      item.Category.Name,
			//FirstImageSrc: item.Category.FirstImageSrc,
		}
		i.Variants = slice.Filter(variants, func(index int, v sVariant.VariantToOrderOut) bool {
			return v.ProductID == item.ID
		})
		return i
	})
	return out, err
}
