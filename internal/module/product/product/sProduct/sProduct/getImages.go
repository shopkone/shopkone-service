package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/setting/file/mFile"
	"shopkone-service/internal/module/setting/file/sFile"
)

type GetProductImagesOut struct {
	ProductId uint
	Files     []string
}

// 获取商品图片
func (s *sProduct) GetProductImages(productIds []uint) (out []GetProductImagesOut, err error) {
	productIds = slice.Unique(productIds)
	productIds = slice.Filter(productIds, func(index int, item uint) bool {
		return item > 0
	})
	var productFileIds []mProduct.ProductFiles
	if err = s.orm.Model(&productFileIds).Where("shop_id = ? AND product_id IN ?", s.shopId, productIds).
		Select("id", "product_id", "file_id").Find(&productFileIds).Error; err != nil {
		return nil, err
	}
	fileIds := slice.Map(productFileIds, func(index int, item mProduct.ProductFiles) uint {
		return item.FileId
	})
	fileIds = slice.Unique(fileIds)
	fileIds = slice.Filter(fileIds, func(index int, item uint) bool {
		return item != 0
	})
	files, err := sFile.NewFile(s.orm, s.shopId).FileListByIds(fileIds)
	if err != nil {
		return nil, err
	}
	out = slice.Map(productIds, func(index int, productId uint) GetProductImagesOut {
		i := GetProductImagesOut{}
		currentProductFileIds := slice.Filter(productFileIds, func(index int, cpf mProduct.ProductFiles) bool {
			return cpf.ProductId == productId
		})
		currentProductFiles := slice.Map(currentProductFileIds, func(index int, cpf mProduct.ProductFiles) vo.FileListByIdsRes {
			find, ok := slice.FindBy(files, func(index int, f vo.FileListByIdsRes) bool {
				return f.Id == cpf.FileId
			})
			if ok {
				return find
			}
			return vo.FileListByIdsRes{}
		})
		currentProductFiles = slice.Filter(currentProductFiles, func(index int, item vo.FileListByIdsRes) bool {
			return item.Type == mFile.FileTypeImage || item.Type == mFile.FileTypeVideo
		})
		i.ProductId = productId
		i.Files = slice.Map(currentProductFiles, func(index int, item vo.FileListByIdsRes) string {
			if item.Type == mFile.FileTypeImage {
				return item.Path
			}
			if item.Type == mFile.FileTypeVideo {
				return item.Cover
			}
			return ""
		})
		i.Files = slice.Filter(i.Files, func(index int, item string) bool {
			return item != ""
		})
		return i
	})
	out = slice.Filter(out, func(index int, item GetProductImagesOut) bool {
		return len(item.Files) > 0
	})
	return out, err
}
