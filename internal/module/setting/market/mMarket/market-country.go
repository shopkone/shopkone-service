package mMarket

import "shopkone-service/internal/module/base/orm/mOrm"

type MarketCountry struct {
	mOrm.Model
	MarketID    uint   `gorm:"not null;uniqueIndex:id_country_code"`
	CountryCode string `gorm:"size:3;uniqueIndex:id_country_code"`
}
