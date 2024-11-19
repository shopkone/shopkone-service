package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/resource/iResource"
	"shopkone-service/internal/module/base/resource/sResource"
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarket) MarketOptions(getT func(string) string) (res []vo.MarketOptionsRes, err error) {
	var list []mMarket.Market
	countries := sResource.NewCountry().List()
	if err = s.orm.Model(&mMarket.Market{}).Where("shop_id = ?", s.shopId).
		Select("id",
			"name",
			"is_main",
			"language_ids",
			"default_language_id",
			"domain_type",
		).
		Find(&list).Error; err != nil {
		return nil, err
	}
	res = slice.Map(list, func(index int, item mMarket.Market) vo.MarketOptionsRes {
		i := vo.MarketOptionsRes{
			Label:             item.Name,
			Value:             item.ID,
			IsMain:            item.IsMain,
			LanguageIds:       item.LanguageIds,
			DefaultLanguageId: item.DefaultLanguageID,
			DomainType:        item.DomainType,
		}
		if i.IsMain && getT != nil {
			country, ok := slice.FindBy(countries, func(index int, country iResource.CountryListOut) bool {
				return country.Code == item.Name
			})
			if ok {
				i.Label = getT(country.Name)
			}
		}
		if i.LanguageIds == nil {
			i.LanguageIds = []uint{}
		}
		return i
	})
	return res, err
}
