package mDomains

import "shopkone-service/internal/module/base/orm/mOrm"

type DomainBlockCountry struct {
	mOrm.Model
	CountryCode string `gorm:"index;not null"`
}
