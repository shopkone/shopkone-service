package mLanguage

import "shopkone-service/internal/module/base/orm/mOrm"

type Language struct {
	mOrm.Model
	Code      string `json:"code"`                    // 语言编码
	IsDefault bool   `json:"is_default" gorm:"index"` // 是否默认
}
