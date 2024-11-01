package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/iProduct"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/setting/file/sFile"
)

func (s *sProduct) ListByIdsWithoutVariants(ids []uint) (res []iProduct.ListByIdsWithoutVariantsOut, err error) {
	// 获取商品列表
	var products []mProduct.Product
	query := s.orm.Model(&products).Where("shop_id = ?", s.shopId)
	query = query.Where("id in (?)", ids)
	query = query.Select("id", "title")
	if err = query.Find(&products).Error; err != nil {
		return res, err
	}
	// 获取图片
	var productImages []mProduct.ProductFiles
	if err = s.orm.Model(&productImages).
		Where("shop_id = ? AND product_id in (?)", s.shopId, ids).
		Select("file_id", "product_id").Find(&productImages).Error; err != nil {
		return res, err
	}
	imageIds := slice.Map(productImages, func(index int, item mProduct.ProductFiles) uint {
		return item.FileId
	})
	files, err := sFile.NewFile(s.orm, s.shopId).FileListByIds(imageIds)
	if err != nil {
		return res, err
	}
	// 组装数据
	res = slice.Map(products, func(index int, item mProduct.Product) iProduct.ListByIdsWithoutVariantsOut {
		temp := iProduct.ListByIdsWithoutVariantsOut{}
		temp.Id = item.ID
		temp.Title = item.Title
		image, has := slice.FindBy(productImages, func(index int, i mProduct.ProductFiles) bool {
			return item.ID == i.ProductId
		})
		if !has {
			return temp
		}
		find, ok := slice.FindBy(files, func(index int, f vo.FileListByIdsRes) bool {
			return f.Id == image.FileId
		})
		if ok {
			temp.Image = find.Path
		}
		return temp
	})
	return res, err
}
