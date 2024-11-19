package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/inventory/iInventory"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"
	"shopkone-service/utility/handle"
)

func (s *sProduct) List(in vo.ProductListReq) (res handle.PageRes[vo.ProductListRes], err error) {
	query := s.orm.Model(&mProduct.Product{}).Where("shop_id = ?", s.shopId)
	// 获取全部总数
	res.Page = in.PageReq
	if err = query.Count(&res.Page.AllTotal).Error; err != nil {
		return res, err
	}
	// 过滤是否追踪库存
	if in.TrackInventory != 0 {
		track := true
		if in.TrackInventory == 2 {
			track = false
		}
		query = query.Where("inventory_tracking = ?", track)
	}
	// 排除和包含
	if in.ExcludeIds != nil && len(*in.ExcludeIds) > 0 {
		query = query.Where("id not in (?)", *in.ExcludeIds)
	}
	if in.IncludeIds != nil {
		query = query.Where("id in (?)", *in.IncludeIds)
	}
	// 获取总数
	if err = query.Count(&res.Total).Error; err != nil {
		return res, err
	}
	// 分页
	query = query.Scopes(handle.Pagination(in.PageReq)).Order("id desc")
	// 获取商品列表
	var products []mProduct.Product
	err = query.Find(&products).Error
	if err != nil {
		return res, err
	}
	// 获取变体列表
	variantService := sVariant.NewVariant(s.orm, s.shopId)
	productIds := slice.Map(products, func(index int, item mProduct.Product) uint {
		return item.ID
	})
	variants, err := variantService.ListByProductIds(productIds)
	// 获取变体库存列表
	variantIds := slice.Map(variants, func(index int, item mProduct.Variant) uint {
		return item.ID
	})
	inventoryQuantity, err := sInventory.NewInventory(s.orm, s.shopId).CountByVariantIds(variantIds, 0)
	if err != nil {
		return res, err
	}
	// 获取变体图片
	variantGetImageIn := slice.Map(variants, func(index int, item mProduct.Variant) sVariant.GetImagesIn {
		i := sVariant.GetImagesIn{}
		i.VariantId = item.ID
		i.ImageId = item.ImageId
		return i
	})
	variantImages, err := variantService.GetImages(variantGetImageIn)
	// 获取商品文件
	productFiles, err := s.GetProductImages(productIds)
	// 组装数据
	res.List = slice.Map(products, func(index int, item mProduct.Product) vo.ProductListRes {
		temp := vo.ProductListRes{}
		temp.Id = item.ID
		temp.CreatedAt = handle.ToUnix(item.CreatedAt)
		temp.Status = item.Status
		temp.Vendor = item.Vendor
		temp.Title = item.Title
		temp.Spu = item.Spu
		temp.InventoryTracking = item.InventoryTracking
		// 组装变体
		productVariants := slice.Filter(variants, func(index int, v mProduct.Variant) bool {
			return v.ProductId == item.ID
		})
		temp.Variants = slice.Map(productVariants, func(index int, v mProduct.Variant) vo.VariantList {
			inventory, _ := slice.FindBy(inventoryQuantity, func(index int, i iInventory.CountByVariantIdsOut) bool {
				return i.VariantId == v.ID
			})
			return vo.VariantList{
				Id:       v.ID,
				Price:    v.Price,
				Sku:      v.Sku,
				Quantity: inventory.Quantity,
				Name:     v.Name,
			}
		})
		// 组装图片
		image, has := slice.FindBy(productFiles, func(index int, i GetProductImagesOut) bool {
			return item.ID == i.ProductId
		})
		if has {
			temp.Image = image.Files[0]
		}
		currentVariantImages := slice.Filter(variantImages, func(index int, cv sVariant.GetImagesOut) bool {
			_, ok := slice.FindBy(temp.Variants, func(index int, v vo.VariantList) bool {
				return v.Id == cv.VariantId
			})
			return ok
		})
		if temp.Image == "" && len(currentVariantImages) > 0 {
			temp.Image = currentVariantImages[0].Image
		}
		return temp
	})
	return res, err
}
