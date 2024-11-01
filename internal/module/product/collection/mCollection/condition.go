package mCollection

import "shopkone-service/internal/module/base/orm/mOrm"

// ProductCondition 商品集合自动匹配条件
type ProductCondition struct {
	mOrm.Model
	Action       string `gorm:"size:100;index;not null"` // 动作
	Key          string `gorm:"size:100;index;not null"` // 键
	Value        string `gorm:"size:250;index"`
	CollectionId uint   `gorm:"index;not null"` // 商品集合Id
}
