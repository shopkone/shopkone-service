package mFile

import "shopkone-service/internal/module/base/orm/mOrm"

type FileGroup struct {
	mOrm.Model
	Name string `gorm:"size:500;not null;index"`
}
