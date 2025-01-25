package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/seo/sSeo"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/product/product/sProduct/sTransfer"
	"shopkone-service/internal/module/product/product/sProduct/sVariant"
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
	var labelImages []mProduct.ProductLabelImage
	if err = s.orm.
		Where("product_id = ? AND shop_id = ?", id, s.shopId).
		Find(&labelImages).Error; err != nil {
		return vo.ProductInfoRes{}, err
	}
	images := slice.Map(labelImages, func(index int, item mProduct.ProductLabelImage) vo.LabelImage {
		return vo.LabelImage{
			ImageId: item.ImageId,
			Label:   item.Label,
			Value:   item.Value,
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
	res.Variants = variants
	res.LabelImages = images
	res.FileIds = fileIds
	return res, err
}
