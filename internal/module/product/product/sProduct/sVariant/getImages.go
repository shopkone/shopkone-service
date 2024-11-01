package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/setting/file/sFile"
)

type GetImagesOut struct {
	VariantId uint
	Image     string
}

func (s *sVariant) GetImages(variants []mProduct.Variant) (out []GetImagesOut, err error) {
	imageIds := slice.Map(variants, func(index int, item mProduct.Variant) uint {
		return item.ImageId
	})
	imageIds = slice.Unique(imageIds)
	imageIds = slice.Filter(imageIds, func(index int, item uint) bool {
		return item > 0
	})
	filesSrc, err := sFile.NewFile(s.orm, s.shopId).FileListByIds(imageIds)
	if err != nil {
		return nil, err
	}
	out = slice.Map(variants, func(index int, item mProduct.Variant) GetImagesOut {
		i := GetImagesOut{VariantId: item.ID}
		find, ok := slice.FindBy(filesSrc, func(index int, v vo.FileListByIdsRes) bool {
			return v.Id == item.ImageId
		})
		if ok {
			i.Image = find.Path
		}
		return i
	})
	out = slice.Filter(out, func(index int, item GetImagesOut) bool {
		return item.Image != ""
	})
	return out, err
}
