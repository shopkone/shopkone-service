package sVariant

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/iProduct"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/internal/module/setting/file/sFile"
)

func (s *sVariant) ListByIds(ids []uint, hasDeleted bool) (res []iProduct.VariantListByIdOut, err error) {
	var variants []mProduct.Variant

	// 获取列表
	query := s.orm.Model(&variants).Where("shop_id = ?", s.shopId)
	query = query.Where("id in (?)", ids)
	if hasDeleted {
		query = query.Unscoped()
	}
	query = query.Select("sku", "id", "price", "image_id", "product_id", "name", "deleted_at")
	if err = query.Find(&variants).Error; err != nil {
		return nil, err
	}

	// 获取图片
	imageIds := slice.Map(variants, func(index int, item mProduct.Variant) uint {
		return item.ImageId
	})
	files, err := sFile.NewFile(s.orm, s.shopId).FileListByIds(imageIds)
	if err != nil {
		return nil, err
	}

	// 转换格式
	res = slice.Map(variants, func(index int, item mProduct.Variant) iProduct.VariantListByIdOut {
		nameArr := slice.Map(item.Name, func(index int, item mProduct.VariantName) string {
			return item.Value
		})
		name := slice.Join(nameArr, " · ")

		image, ok := slice.FindBy(files, func(index int, f vo.FileListByIdsRes) bool {
			return item.ImageId == f.Id
		})

		i := iProduct.VariantListByIdOut{}
		i.Id = item.ID
		i.Name = name
		i.Price = item.Price
		if ok {
			i.Image = image.Path
		}
		i.Sku = item.Sku
		i.ProductId = item.ProductId
		if item.DeletedAt.Valid {
			i.IsDeleted = true
		}
		return i
	})

	return res, nil
}
