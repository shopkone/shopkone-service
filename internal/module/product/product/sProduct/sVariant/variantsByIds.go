package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/product/product/mProduct"
)

type VariantByIdsIn struct {
	VariantIds   []uint
	HasDeleted   bool
	HasInventory bool
	LocationId   uint
}

type VariantByIdsOut struct {
	Id             uint
	Image          string
	Name           []mProduct.VariantName
	Price          float32
	Inventory      uint
	TrackInventory bool
	ProductId      uint
	IsDeleted      bool
}

func (s *sVariant) VariantsByIds(in VariantByIdsIn) (out []VariantByIdsOut, err error) {
	var variants []mProduct.Variant
	query := s.orm.Model(&mProduct.Variant{}).Where("id in (?)", in.VariantIds)

	// 是否包含删除的
	if in.HasDeleted {
		query = query.Unscoped()
	}
	// 查询变体
	query = query.Select("id", "image_id", "name", "price", "product_id")
	if err = query.Find(&variants).Error; err != nil {
		return nil, err
	}
	// 查询库存
	var inventories []mInventory.Inventory
	if in.HasInventory {
		inventories, err = sInventory.NewInventory(s.orm, s.shopId).ListByVariantsIds(in.VariantIds, in.LocationId)
		if err != nil {
			return nil, err
		}
	}
	// 获取图片
	variantGetImageIn := slice.Map(variants, func(index int, item mProduct.Variant) GetImagesIn {
		i := GetImagesIn{}
		i.VariantId = item.ID
		i.ImageId = item.ImageId
		return i
	})
	images, err := s.GetImages(variantGetImageIn)
	if err != nil {
		return nil, err
	}
	// 组装数据
	out = slice.Map(variants, func(index int, item mProduct.Variant) VariantByIdsOut {
		i := VariantByIdsOut{}
		i.Id = item.ID
		i.Name = item.Name
		i.Price = item.Price
		i.ProductId = item.ProductId
		// 获取图片
		image, ok := slice.FindBy(images, func(index int, image GetImagesOut) bool {
			return image.Image != "" && image.VariantId == item.ID
		})
		if ok {
			i.Image = image.Image
		}
		// 获取库存
		if in.HasInventory {
			slice.ForEach(inventories, func(index int, inventory mInventory.Inventory) {
				if inventory.VariantId == item.ID {
					i.Inventory = inventory.Quantity + i.Inventory
				}
			})
		}
		return i
	})
	return out, err
}
