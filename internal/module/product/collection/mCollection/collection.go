package mCollection

import (
	"shopkone-service/internal/module/base/orm/mOrm"
)

type CollectionType uint8 // 匹配模式

const (
	CollectionTypeManual CollectionType = iota + 1 // 手动模式
	CollectionTypeAuto                             // 自动模式
)

type CollectionMatchMode uint8 // 自动匹配模式

const (
	CollectionMatchModeAll CollectionMatchMode = iota + 1 // 满足所有条件
	CollectionMatchModeAny                                // 满足任一条件
)

/*ProductCollection 商品专辑*/
type ProductCollection struct {
	mOrm.Model
	Title          string              `gorm:"size:200"`       // 标题
	Description    string              `gorm:"size:5000"`      // 描述
	CollectionType CollectionType      `gorm:"default:1"`      // 手动/自动模式
	MatchMode      CollectionMatchMode `gorm:"default:1"`      // 自动下的匹配模式
	SeoId          uint                `gorm:"index;not null"` // seo id
	CoverId        uint                `gorm:"default:null"`   // 封面图id
	IsAllType      bool                `gorm:"default:false"`  // 是否包含所有
}

type CollectionProduct struct {
	mOrm.Model
	CollectionId uint `gorm:"index;not null"`
	ProductId    uint `gorm:"index;not null"`
}
