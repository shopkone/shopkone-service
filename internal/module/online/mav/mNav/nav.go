package mNav

import "shopkone-service/internal/module/base/orm/mOrm"

type Nav struct {
	mOrm.Model
	Handle string
	Title  string `gorm:"index;not null"`
}

type NavItem struct {
	mOrm.Model
	Levels   uint8
	Title    string
	Url      string
	NavId    uint `gorm:"index;not null"`
	ParentID uint
}
