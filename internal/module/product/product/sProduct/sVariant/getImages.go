package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/file/sFile"
)

type GetImagesIn struct {
	VariantId uint
	ImageId   uint
}

type GetImagesOut struct {
	VariantId uint
	Image     string
}

func (s *sVariant) GetImages(in []GetImagesIn) (out []GetImagesOut, err error) {
	imageIds := slice.Map(in, func(index int, item GetImagesIn) uint {
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
	out = slice.Map(in, func(index int, item GetImagesIn) GetImagesOut {
		i := GetImagesOut{VariantId: item.VariantId}
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
