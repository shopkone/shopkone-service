package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"
)

func (s *sProduct) VariantsWithProduct(in *vo.VariantListByIdsReq) (out []vo.VariantListByIdsRes, err error) {
	// 获取list
	variantService := sVariant.NewVariant(s.orm, s.shopId)
	variantsIn := sVariant.VariantByIdsIn{
		VariantIds:   in.Ids,
		HasDeleted:   in.HasDeleted,
		HasInventory: in.HasInventory,
		LocationId:   in.LocationID,
	}
	variants, err := variantService.VariantsByIds(variantsIn)
	if err != nil {
		return nil, err
	}
	// 获取商品
	productIds := slice.Map(variants, func(index int, item sVariant.VariantByIdsOut) uint {
		return item.ProductId
	})
	productIds = slice.Unique(productIds)
	products, err := s.ProductOptions(productIds)
	if err != nil {
		return nil, err
	}
	// 组装数据
	out = slice.Map(variants, func(index int, item sVariant.VariantByIdsOut) vo.VariantListByIdsRes {
		product, ok := slice.FindBy(products, func(index int, product ProductOptionsOut) bool {
			return product.Id == item.ProductId
		})
		i := vo.VariantListByIdsRes{
			Id:                item.Id,
			Image:             item.Image,
			Name:              variantService.VariantNameToString(item.Name),
			ProductTitle:      "",
			IsDeleted:         false,
			Price:             item.Price,
			Inventory:         item.Inventory,
			InventoryTracking: false,
			InventoryPolicy:   0,
		}
		if ok {
			i.ProductTitle = product.Title
			i.InventoryPolicy = product.InventoryPolicy
			i.InventoryTracking = product.InventoryTracking
			if i.Image == "" && len(product.Images) > 0 {
				i.Image = product.Images[0]
			}
		}
		return i
	})
	return out, err
}
