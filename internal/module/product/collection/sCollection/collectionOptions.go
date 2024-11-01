package sCollection

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/collection/mCollection"
)

func (s *sCollection) CollectionOptions() (options []vo.CollectionOptionsRes, err error) {
	var list []mCollection.ProductCollection
	query := s.orm.Model(&list)
	query = query.Where("shop_id = ?", s.shopId)
	query = query.Select("id", "collection_type", "title")
	if err = query.Find(&list).Error; err != nil {
		return options, err
	}

	options = slice.Map(list, func(index int, item mCollection.ProductCollection) vo.CollectionOptionsRes {
		i := vo.CollectionOptionsRes{}
		i.Value = item.ID
		i.Label = item.Title
		if item.CollectionType == mCollection.CollectionTypeAuto {
			i.Disabled = true
		}
		return i
	})

	return options, err
}
