package sProduct

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/product/mProduct"
	"shopkone-service/utility/code"
)

func (s *sProduct) ListByIds(ids []uint) (res []vo.ListByIdsRes, err error) {
	if len(ids) > 500 {
		return res, code.OnceMaxFind500
	}
	// 获取商品列表
	var products []mProduct.Product
	query := s.orm.Model(&mProduct.Product{}).Where("shop_id = ?", s.shopId)
	query = query.Where("id IN (?)", ids)
	query = query.Select("id", "title", "status")
	if err = query.Find(&products).Error; err != nil {
		return res, err
	}

	// 获取文件fileIds
	productIds := slice.Map(products, func(index int, item mProduct.Product) uint {
		return item.ID
	})
	productFiles, err := s.GetProductImages(productIds)
	if err != nil {
		return nil, err
	}
	res = slice.Map(products, func(index int, item mProduct.Product) vo.ListByIdsRes {
		temp := vo.ListByIdsRes{}
		temp.Id = item.ID
		temp.Status = item.Status
		temp.Title = item.Title
		// 组装图片
		image, has := slice.FindBy(productFiles, func(index int, i GetProductImagesOut) bool {
			return item.ID == i.ProductId
		})
		if has {
			temp.Image = image.Files[0]
		}
		return temp
	})
	return res, err
}
