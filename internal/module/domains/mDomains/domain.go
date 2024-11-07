package mDomains

import "shopkone-service/internal/module/base/orm/mOrm"

type Domain struct {
	mOrm.Model
	IsMain    bool   `gorm:"index"`
	IsDefault bool   `json:"is_default"`
	Domain    string `gorm:"size:500"`
	Type      uint8  `gorm:"default:0"`
	Ip        string `gorm:"index"`
}
