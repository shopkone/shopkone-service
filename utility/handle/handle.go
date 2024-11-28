package handle

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"shopkone-service/internal/consts"
)

func GenUid() string {
	uid := uuid.New().String()
	// 取出uuid中的-
	uid = uid[0:8] + uid[9:13] + uid[14:18] + uid[19:23] + uid[24:]
	return uid + guid.S()
}

type IDGetter interface {
	GetID() uint
}

func DiffUpdate[T IDGetter](news, olds []T) (insert, update, remove []T, err error) {
	newIDs := []uint{}
	oldIDs := []uint{}
	for _, item := range news {
		newIDs = append(newIDs, item.GetID())
	}
	for _, item := range olds {
		oldIDs = append(oldIDs, item.GetID())
	}
	InsertIDs := slice.Filter[uint](newIDs, func(index int, item uint) bool {
		return !slice.Contain[uint](oldIDs, item)
	})
	UpdateIDs := slice.Filter[uint](newIDs, func(index int, item uint) bool {
		return slice.Contain[uint](oldIDs, item)
	})
	RemoveIDs := slice.Filter[uint](oldIDs, func(index int, item uint) bool {
		return !slice.Contain[uint](newIDs, item)
	})
	insert = slice.Filter(news, func(index int, item T) bool {
		return slice.Contain[uint](InsertIDs, item.GetID())
	})
	update = slice.Filter(news, func(index int, item T) bool {
		return slice.Contain[uint](UpdateIDs, item.GetID())
	})
	remove = slice.Filter(olds, func(index int, item T) bool {
		return slice.Contain[uint](RemoveIDs, item.GetID())
	})
	return insert, update, remove, err
}

type BatchUpdateByIdIn struct {
	Orm    *gorm.DB // 数据库连接
	ShopID uint     // 店铺ID
	Query  []string // 更新字段
}

func BatchUpdateById[T any](in BatchUpdateByIdIn, list *[]T) (err error) {
	if len(*list) == 0 {
		return nil
	}
	query := in.Orm.Model(list).Where("shop_id = ?", in.ShopID)
	conflict := clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(in.Query),
	}
	return query.Clauses(conflict).Create(&list).Error
}

func RoundMoney(money float64) float64 {
	return float64(int(money*100)) / 100
}

func RoundMoney32(money float32) float32 {
	return float32(int(money*100)) / 100
}

func ToKg(weight float32, uint consts.WeightUnit) float32 {
	if uint == consts.WeightUnitGram {
		return weight / 1000
	}
	if uint == consts.WeightUnitOunce {
		return weight / 35.2739619
	}
	if uint == consts.WeightUnitPound {
		return weight / 2.20462262
	}
	if uint == consts.WeightUnitKilogram {
		return weight
	}
	return weight
}
