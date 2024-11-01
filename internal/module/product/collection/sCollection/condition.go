package sCollection

import (
	"github.com/duke-git/lancet/v2/slice"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/collection/iCollection"
	"shopkone-service/internal/module/product/collection/mCollection"
	"shopkone-service/utility/code"
)

type sCondition struct {
	shopId uint
	orm    *gorm.DB
}

func NewCondition(orm *gorm.DB, shopId uint) *sCondition {
	return &sCondition{shopId, orm}
}

type ICondition interface {
	// CreateList 创建集合条件
	CreateList(in iCollection.CreateListIn) (err error)
	// MatchProducts 匹配商品
	MatchProducts(in iCollection.MatchProductsIn) (err error)
	// UpdateList 更新条件
	UpdateList(in iCollection.UpdateListIn) (err error)
	// ListByCollectionId 根据集合ID获取条件列表
	ListByCollectionId(collectionId uint) (res []mCollection.ProductCondition, err error)
	// IsChange 判断是否需要更新
	IsChange(newConditions, oldConditions []vo.BaseCondition) bool
}

func (s *sCondition) CreateList(in iCollection.CreateListIn) (err error) {
	// 校验key
	validKey := slice.Every(in.Conditions, func(index int, item vo.BaseCondition) bool {
		return FilterKey(item.Key)
	})
	if !validKey {
		return code.InValidKey
	}

	// 校验操作
	validAction := slice.Every(in.Conditions, func(index int, item vo.BaseCondition) bool {
		return FilterAction(item.Action)
	})
	if !validAction {
		return code.InValidAction
	}

	// 创建
	list := slice.Map(in.Conditions, func(index int, item vo.BaseCondition) mCollection.ProductCondition {
		i := mCollection.ProductCondition{
			CollectionId: in.CollectionId,
			Action:       item.Action,
			Key:          item.Key,
			Value:        item.Value,
		}
		i.ShopId = s.shopId
		return i
	})
	if err = s.orm.Create(list).Error; err != nil {
		return err
	}

	// 匹配商品
	matchIn := iCollection.MatchProductsIn{
		CollectionId: in.CollectionId,
		Conditions:   in.Conditions,
		MatchMode:    in.MatchMode,
	}
	return s.MatchProducts(matchIn)
}

func (s *sCondition) UpdateList(in iCollection.UpdateListIn) (err error) {
	// 删除之前的条件
	query := s.orm.Model(&mCollection.ProductCondition{}).Where("shop_id = ? AND collection_id = ?", s.shopId, in.CollectionId)
	if err = query.Unscoped().Delete(&mCollection.ProductCondition{}).Error; err != nil {
		return err
	}
	// 创建新的条件
	createIn := iCollection.CreateListIn{
		CollectionId: in.CollectionId,
		Conditions:   in.Conditions,
		MatchMode:    in.MatchMode,
	}
	return s.CreateList(createIn)
}

func (s *sCondition) ListByCollectionId(collectionId uint) (res []mCollection.ProductCondition, err error) {
	query := s.orm.Model(&mCollection.ProductCondition{})
	query = query.Where("shop_id = ? AND collection_id = ?", s.shopId, collectionId)
	query = query.Select("id", "action", "key", "value")
	return res, query.Find(&res).Error
}

// IsChange 检查新旧条件是否有变化
func (s *sCondition) IsChange(newConditions, oldConditions []mCollection.ProductCondition) bool {

	if len(newConditions) != len(oldConditions) {
		return true
	}

	return !slice.Every(newConditions, func(index int, nc mCollection.ProductCondition) bool {
		_, ok := slice.FindBy(oldConditions, func(index int, item mCollection.ProductCondition) bool {
			return item.Action == nc.Action && item.Key == nc.Key && item.Value == nc.Value
		})
		return ok
	})
}
