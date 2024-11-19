package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/product/mProduct"
)

type ProductOptionsOut struct {
	Id                uint
	Images            []string
	Title             string
	InventoryTracking bool
	InventoryPolicy   mProduct.InventoryPolicy
}

func (s *sProduct) ProductOptions(ids []uint) (out []ProductOptionsOut, err error) {
	// 获取商品列表
	var products []mProduct.Product
	query := s.orm.Model(&products).Where("shop_id = ?", s.shopId)
	query = query.Where("id IN ?", ids)
	query = query.Select("id", "title", "inventory_tracking", "inventory_policy")
	if err = query.Find(&products).Error; err != nil {
		return out, err
	}

	// 获取图片
	images, err := s.GetProductImages(ids)
	if err != nil {
		return nil, err
	}
	out = slice.Map(ids, func(index int, id uint) ProductOptionsOut {
		i := ProductOptionsOut{}
		product, ok := slice.FindBy(products, func(index int, product mProduct.Product) bool {
			return product.ID == id
		})
		if ok {
			i.Id = product.ID
			i.Title = product.Title
			i.InventoryPolicy = product.InventoryPolicy
			i.InventoryTracking = product.InventoryTracking
		}
		slice.ForEach(images, func(index int, item GetProductImagesOut) {
			if item.ProductId == id {
				i.Images = slice.Concat(i.Images, item.Files)
			}
		})
		return i
	})
	return out, err
}
