package mLanguage

import "shopkone-service/internal/module/base/orm/mOrm"

type Language struct {
	mOrm.Model
	Code      string `json:"code"`                    // 语言编码
	IsActive  bool   `json:"is_active" gorm:"index"`  // 是否启用
	IsDefault bool   `json:"is_default" gorm:"index"` // 是否默认
}
