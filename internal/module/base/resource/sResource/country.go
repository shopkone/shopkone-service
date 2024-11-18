package sResource

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/base/resource"
	"shopkone-service/internal/module/base/resource/iResource"
	"shopkone-service/internal/module/base/resource/mResource"
	"shopkone-service/utility/code"
)

type ICountry interface {
	// List 获取国家列表
	List() []iResource.CountryListOut
	// CountryByCode 获取国家信息
	CountryByCode(c string) (mResource.Country, error)
	// ZoneByCode 获取省份信息
	ZoneByCode(c string) (mResource.CountryZone, error)
	// PhonePrefixList 获取国家电话前缀列表
	PhonePrefixList() []iResource.PhonePrefixListOut
	// PhonePrefixByCountryCode 获取国家电话前缀
	PhonePrefixByCountryCode(code string) iResource.PhonePrefixListOut
}

type sCountry struct{}

func NewCountry() *sCountry {
	return &sCountry{}
}

func (s *sCountry) List() []iResource.CountryListOut {
	result := slice.Map(resource.Countries, func(index int, item mResource.Country) iResource.CountryListOut {
		c := iResource.CountryListOut{}
		c.Code = item.Code
		c.Name = item.Name
		c.Continent = item.Continent
		c.Flag = item.Flag
		c.Formatting = item.AddressFormatting
		c.Zones = slice.Map(item.Zones, func(index int, zone mResource.CountryZone) iResource.ZoneListOut {
			return iResource.ZoneListOut{Code: zone.Code, Name: zone.Name}
		})
		c.Config = item.AddressLabel
		c.PostalCodeConfig = item.PostalCodeConfig
		return c
	})
	return result
}

func (s *sCountry) CountryByCode(c string) (mResource.Country, error) {
	result, ok := slice.FindBy(resource.Countries, func(index int, item mResource.Country) bool {
		return item.Code == c
	})
	if !ok {
		return mResource.Country{}, code.CountryCodeUnknown
	}
	return result, nil
}

func (s *sCountry) ZoneByCode(c string) (mResource.CountryZone, error) {
	if c == "" {
		return mResource.CountryZone{}, nil
	}
	result := mResource.CountryZone{}
	slice.ForEach(resource.Countries, func(index int, item mResource.Country) {
		slice.ForEach(item.Zones, func(index int, zone mResource.CountryZone) {
			if zone.Code == c {
				result = zone
			}
		})
	})
	if result.Code == "" {
		return result, code.ZoneCodeUnknown
	}
	return result, nil
}

func (s *sCountry) PhonePrefixList() []iResource.PhonePrefixListOut {
	result := slice.Map(resource.Countries, func(index int, item mResource.Country) iResource.PhonePrefixListOut {
		return iResource.PhonePrefixListOut{Code: item.Code, Prefix: item.PhoneNumberPrefix}
	})
	return result
}

func (s *sCountry) PhonePrefixByCountryCode(code string) iResource.PhonePrefixListOut {
	list := s.PhonePrefixList()
	result, _ := slice.FindBy(list, func(index int, item iResource.PhonePrefixListOut) bool {
		return item.Code == code
	})
	return result
}

func (s *sCountry) CheckCountryListExist(countryCodes []string) (err error) {
	list := s.List()
	isAllExist := slice.Every(countryCodes, func(index int, code string) bool {
		_, ok := slice.FindBy(list, func(index int, item iResource.CountryListOut) bool {
			return item.Code == code
		})
		return ok
	})
	if !isAllExist {
		err = code.CountryCodeUnknown
	}
	return
}
