package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/product/product/mProduct"
)

func (s *sProduct) UpdateProductFiles(productId uint, fileIds []uint) (err error) {
	// 删除旧的商品文件
	if err = s.orm.Where("product_id = ? AND shop_id = ?", productId, s.shopId).
		Delete(&mProduct.ProductFiles{}).Error; err != nil {
		return err
	}
	// 创建新的商品文件
	files := slice.Map(fileIds, func(index int, item uint) mProduct.ProductFiles {
		temp := mProduct.ProductFiles{
			ProductId: productId,
			FileId:    item,
			Position:  uint(index),
		}
		temp.ShopId = s.shopId
		return temp
	})
	if len(files) > 0 {
		if err = s.orm.Create(&files).Error; err != nil {
			return err
		}
	}
	return err
}
