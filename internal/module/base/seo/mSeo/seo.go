package mSeo

import "shopkone-service/internal/module/base/orm/mOrm"

type Seo struct {
	mOrm.Model
	MetaDescription string `json:"meta_description"  gorm:"type:text"`
	PageTitle       string `json:"page_title"  v:"required" gorm:"size:200"`
	UrlHandle       string `json:"url_handle" gorm:"size:200"`
	Follow          bool   `json:"follow"`
}
