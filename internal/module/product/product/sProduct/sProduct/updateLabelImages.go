package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/mProduct"
)

func (s *sProduct) UpdateLabelImages(productId uint, labelImages []vo.LabelImage) (err error) {
	if len(labelImages) == 0 {
		return nil
	}
	// 删除旧的商品标签图片
	if err = s.orm.Where("product_id = ? AND shop_id = ?", productId, s.shopId).
		Delete(&mProduct.ProductLabelImage{}).Error; err != nil {
		return err
	}
	// 创建新的商品标签图片
	images := slice.Map(labelImages, func(index int, item vo.LabelImage) mProduct.ProductLabelImage {
		i := mProduct.ProductLabelImage{}
		i.Value = item.Value
		i.Label = item.Label
		i.ImageId = item.ImageId
		i.ShopId = s.shopId
		i.ProductId = productId
		return i
	})
	images = slice.Filter(images, func(index int, item mProduct.ProductLabelImage) bool {
		return item.Label != "" && item.Value != "" && item.ImageId > 0 && index < 500
	})
	if err = s.orm.Create(&images).Error; err != nil {
		return err
	}
	return err
}
