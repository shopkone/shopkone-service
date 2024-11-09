package sMarket

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
)

func (s *sMarket) MarketInfo(id uint) (out vo.MarketInfoRes, err error) {
	var data mMarket.Market
	if err = s.orm.Model(&data).Where("shop_id = ? AND id = ?", s.shopId, id).
		Omit("shop_id", "created_at", "updated_at", "deleted_at").First(&data).
		Error; err != nil {
		return out, err
	}

	countries, err := s.CountryList([]uint{id})
	if err != nil {
		return vo.MarketInfoRes{}, err
	}

	out.ID = data.ID
	out.Name = data.Name
	out.IsMain = data.IsMain
	out.CountryCodes = slice.Map(countries, func(index int, item mMarket.MarketCountry) string {
		return item.CountryCode
	})
	out.DomainType = data.DomainType
	out.DomainSuffix = data.DomainSuffix
	out.SubDomainID = data.SubDomainID
	out.LanguageIds = data.LanguageIds
	out.DefaultLanguageId = data.DefaultLanguageID
	return out, err
}
