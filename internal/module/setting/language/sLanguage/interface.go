package sLanguage

import "gorm.io/gorm"

type sLanguage struct {
	orm    *gorm.DB
	shopId uint
}

func NewLanguage(orm *gorm.DB, shopId uint) *sLanguage {
	return &sLanguage{orm, shopId}
}
