package sProduct

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/seo/sSeo"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/product/product/sProduct/sTransfer"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"

	"github.com/duke-git/lancet/v2/slice"
)

func (s *sProduct) Info(id uint) (res vo.ProductInfoRes, err error) {
	// 获取商品信息
	product := mProduct.Product{}
	err = s.orm.Where("id = ?", id).First(&product).Error
	if err != nil {
		return vo.ProductInfoRes{}, err
	}
	// 获取seo
	seo, err := sSeo.NewSeo(s.orm, s.shopId).Info(product.SeoId)
	if err != nil {
		return vo.ProductInfoRes{}, err
	}
	// 获取变体
	variants, err := sVariant.NewVariant(s.orm, s.shopId).ListByProductId(id)
	if err != nil {
		return vo.ProductInfoRes{}, err
	}
	// 获取图片Ids
	var files []mProduct.ProductFiles
	if err = s.orm.
		Where("product_id = ? AND shop_id = ?", id, s.shopId).
		Select("file_id").Find(&files).Error; err != nil {
		return vo.ProductInfoRes{}, err
	}
	fileIds := slice.Map(files, func(index int, item mProduct.ProductFiles) uint {
		return item.FileId
	})
	// 获取标签图片
	var ProductOptions []mProduct.ProductOption
	if err = s.orm.
		Where("product_id = ? AND shop_id = ?", id, s.shopId).
		Find(&ProductOptions).Error; err != nil {
		return vo.ProductInfoRes{}, err
	}
	productOptions := slice.Map(ProductOptions, func(index int, item mProduct.ProductOption) vo.ProductOption {
		return vo.ProductOption{
			ImageId: item.ImageId,
			Label:   item.Label,
			Values:  item.Values,
			Id:      item.ID,
		}
	})
	// 组装数据
	res.ProductCreateReq = vo.ProductCreateReq{
		BaseProduct: sTransfer.NewProductTransfer(s.shopId).ModelToProduct(product),
	}
	// 获取集合
	if res.Collections, err = s.GetCollections(product.ID); err != nil {
		return res, err
	}
	res.Seo = seo
	res.Variants = slice.Map(product.VariantOrder, func(index int, id uint) vo.BaseVariant {
		variant, _ := slice.FindBy(variants, func(index int, item vo.BaseVariant) bool {
			return id == item.Id
		})
		return variant
	})
	res.Variants = slice.Filter(res.Variants, func(index int, item vo.BaseVariant) bool {
		return item.Id != 0
	})
	res.ProductOptions = productOptions
	res.FileIds = fileIds
	return res, err
}
