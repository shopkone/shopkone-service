package sCollection

import (
	"github.com/duke-git/lancet/v2/slice"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/seo/sSeo"
	"shopkone-service/internal/module/product/collection/iCollection"
	"shopkone-service/internal/module/product/collection/mCollection"
	"shopkone-service/internal/module/setting/file/sFile"
	"shopkone-service/utility/code"
	"shopkone-service/utility/handle"
)

type ICollection interface {
	// Create 创建集合
	Create(in vo.CreateProductCollectionReq) (id uint, err error)
	// Update 更新集合
	Update(in vo.UpdateProductCollectionReq) error
	// Remove 删除集合
	Remove()
	// Info 获取集合详情
	Info(id uint) (res vo.ProductCollectionInfoRes, err error)
	// List 获取集合列表
	List(in vo.ProductCollectionListReq) (res handle.PageRes[vo.ProductCollectionListRes], err error)
	// SimpleList 获取集合列表
	SimpleList()
}

type sCollection struct {
	shopId uint
	orm    *gorm.DB
}

func NewCollection(orm *gorm.DB, shopId uint) *sCollection {
	return &sCollection{shopId, orm}
}

func (s *sCollection) Create(in vo.CreateProductCollectionReq) (id uint, err error) {
	// 创建 seo
	seoId, err := sSeo.NewSeo(s.orm, s.shopId).Create(in.Seo)
	if err != nil {
		return 0, err
	}

	// 创建集合
	data := mCollection.ProductCollection{
		Title:          in.Title,
		Description:    in.Description,
		CollectionType: in.CollectionType,
		MatchMode:      in.MatchMode,
		SeoId:          seoId,
		CoverId:        in.CoverId,
	}
	data.ShopId = s.shopId
	if err = s.orm.Create(&data).Error; err != nil {
		return 0, err
	}

	// 创建条件
	if in.CollectionType == mCollection.CollectionTypeAuto {
		// 如果没有条件，则返回
		if len(in.Conditions) == 0 {
			return id, code.ProductConditionMust
		}
		createIn := iCollection.CreateListIn{
			CollectionId: data.ID,
			Conditions:   in.Conditions,
			MatchMode:    in.MatchMode,
		}
		if err = NewCondition(s.orm, s.shopId).CreateList(createIn); err != nil {
			return 0, err
		}
		// 开始匹配现有商品
	}

	// 关联商品
	if in.CollectionType == mCollection.CollectionTypeManual && len(in.ProductIds) > 0 {
		cp := slice.Map(in.ProductIds, func(index int, item uint) mCollection.CollectionProduct {
			i := mCollection.CollectionProduct{}
			i.ProductId = item
			i.CollectionId = data.ID
			i.ShopId = s.shopId
			return i
		})
		return data.ID, s.orm.Create(&cp).Error
	}

	return data.ID, err
}

func (s *sCollection) List(in vo.ProductCollectionListReq) (res handle.PageRes[vo.ProductCollectionListRes], err error) {
	var list []mCollection.ProductCollection
	query := s.orm.Model(&mCollection.ProductCollection{}).Where("shop_id = ?", s.shopId)

	// 查找总数
	if err = query.Count(&res.Total).Error; err != nil {
		return res, err
	}

	// 查找列表
	query = query.Select("id", "title", "collection_type", "cover_id")
	query = query.Scopes(handle.Pagination(in.PageReq)).Order("id desc")
	if err = query.Find(&list).Error; err != nil {
		return res, err
	}

	// 获取封面
	coverIds := slice.Map(list, func(index int, item mCollection.ProductCollection) uint {
		return item.CoverId
	})
	coverIds = slice.Unique(coverIds)
	files, err := sFile.NewFile(s.orm, s.shopId).FileListByIds(coverIds)
	if err != nil {
		return res, err
	}

	// 获取商品数
	collectionIds := slice.Map(list, func(index int, item mCollection.ProductCollection) uint {
		return item.ID
	})
	var collectionProducts []mCollection.CollectionProduct
	if err = s.orm.Model(&mCollection.CollectionProduct{}).Where("shop_id =?", s.shopId).
		Select("collection_id").
		Where("collection_id IN (?)", collectionIds).Find(&collectionProducts).Error; err != nil {
		return res, err
	}

	// 组装数据
	res.List = slice.Map(list, func(index int, item mCollection.ProductCollection) vo.ProductCollectionListRes {
		i := vo.ProductCollectionListRes{}
		i.Id = item.ID
		i.Title = item.Title
		i.CollectionType = item.CollectionType
		// 组装封面
		cover, ok := slice.FindBy(files, func(index int, f vo.FileListByIdsRes) bool {
			return f.Id == item.CoverId
		})
		if ok {
			i.Cover = cover.Path
		}
		// 组装商品数
		pc := slice.Filter(collectionProducts, func(index int, c mCollection.CollectionProduct) bool {
			return c.CollectionId == item.ID
		})
		i.ProductQuantity = len(pc)
		return i
	})

	return res, err
}

func (s *sCollection) Info(id uint) (res vo.ProductCollectionInfoRes, err error) {
	var data mCollection.ProductCollection
	query := s.orm.Model(&data).Where("shop_id = ?", s.shopId)
	query = query.Where("id = ?", id)
	query = query.Select("id", "title", "collection_type", "cover_id", "description", "match_mode", "seo_id")
	if err = query.First(&data).Error; err != nil {
		return res, err
	}
	// 获取 seo
	seo, err := sSeo.NewSeo(s.orm, s.shopId).Info(data.SeoId)
	// 获取商品列表
	var collectionProducts []mCollection.CollectionProduct
	if err = s.orm.Model(&mCollection.CollectionProduct{}).
		Where("collection_id = ?", data.ID).Select("id", "product_id", "collection_id").
		Find(&collectionProducts).Error; err != nil {
		return res, err
	}
	productIds := slice.Map(collectionProducts, func(index int, item mCollection.CollectionProduct) uint {
		return item.ProductId
	})

	// 如果是自动匹配模式，获取条件
	if data.CollectionType == mCollection.CollectionTypeAuto {
		conditions, err := NewCondition(s.orm, s.shopId).ListByCollectionId(data.ID)
		if err != nil {
			return res, err
		}
		res.Conditions = slice.Map(conditions, func(index int, item mCollection.ProductCondition) vo.BaseCondition {
			i := vo.BaseCondition{
				Action: item.Action,
				Key:    item.Key,
				Value:  item.Value,
			}
			return i
		})
	}

	// 组装数据
	res.Id = data.ID
	res.Title = data.Title
	res.Description = data.Description
	res.CollectionType = data.CollectionType
	res.Seo = seo
	res.ProductIds = productIds
	res.MatchMode = data.MatchMode
	res.CoverId = data.CoverId

	return res, err
}

func (s *sCollection) Update(in vo.UpdateProductCollectionReq) (err error) {
	// 获取集合信息
	var data mCollection.ProductCollection
	query := s.orm.Model(&data).Where("shop_id = ?", s.shopId)
	query = query.Where("id = ?", in.Id).Select("id", "seo_id", "collection_type")
	if err = query.First(&data).Error; err != nil {
		return err
	}

	// 不允许修改集合类型
	if data.CollectionType != in.CollectionType {
		return code.ProductCollectionTypeNotAllow
	}

	// 更新seo
	in.Seo.ID = data.SeoId
	if err = sSeo.NewSeo(s.orm, s.shopId).Update(in.Seo); err != nil {
		return err
	}

	// 更新集合
	data.Title = in.Title
	data.Description = in.Description
	data.CollectionType = in.CollectionType
	data.MatchMode = in.MatchMode
	data.CoverId = in.CoverId
	if err = s.orm.Model(&data).
		Where("shop_id = ? AND id = ?", s.shopId, data.ID).
		Select("title", "description", "match_mode", "cover_id").
		Updates(&data).Error; err != nil {
		return err
	}

	// 如果是自动匹配模式，判断是否有更新，没有更新则不继续了
	if in.CollectionType == mCollection.CollectionTypeAuto {
		sConditions := NewCondition(s.orm, s.shopId)
		oldConditions, err := sConditions.ListByCollectionId(data.ID)
		if err != nil {
			return err
		}
		newConditions := slice.Map(in.Conditions, func(index int, item vo.BaseCondition) mCollection.ProductCondition {
			i := mCollection.ProductCondition{
				Action:       item.Action,
				CollectionId: data.ID,
				Key:          item.Key,
				Value:        item.Value,
			}
			i.ShopId = s.shopId
			i.ID = item.Id
			return i
		})
		if !sConditions.IsChange(newConditions, oldConditions) {
			return err
		}
	}

	// 删除之前的商品关联
	q := s.orm.Model(&mCollection.CollectionProduct{})
	q = q.Where("shop_id = ? AND collection_id = ?", s.shopId, data.ID)
	if err = q.Unscoped().Delete(&mCollection.CollectionProduct{}).Error; err != nil {
		return err
	}

	// 更新商品关联
	if in.CollectionType == mCollection.CollectionTypeManual && len(in.ProductIds) > 0 {
		cps := slice.Map(in.ProductIds, func(index int, item uint) mCollection.CollectionProduct {
			i := mCollection.CollectionProduct{}
			i.ProductId = item
			i.CollectionId = data.ID
			i.ShopId = s.shopId
			return i
		})
		if err = s.orm.Create(&cps).Error; err != nil {
			return err
		}
	}

	if in.CollectionType == mCollection.CollectionTypeAuto {
		// 如果没有条件，则返回
		if len(in.Conditions) == 0 {
			return code.ProductConditionMust
		}
		// 更新条件
		updateIn := iCollection.UpdateListIn{
			CollectionId: data.ID,
			Conditions:   in.Conditions,
			MatchMode:    in.MatchMode,
		}
		return NewCondition(s.orm, s.shopId).UpdateList(updateIn)
	}
	return err
}
