package sResource

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/base/resource"
	"shopkone-service/internal/module/base/resource/mResource"
	"strings"
)

type ITimezone interface {
	// List 获取时区列表
	List() []mResource.Timezone
	// TimezoneByCountry 根据国家获取时区
	TimezoneByCountry(c string) (string, error)
}

type sTimezone struct {
}

func NewTimezone() *sTimezone {
	return &sTimezone{}
}

func (s *sTimezone) List() []mResource.Timezone {
	return resource.Timezones
}

func (s *sTimezone) TimezoneByCountry(c string) (string, error) {
	country, err := NewCountry().CountryByCode(c)
	if err != nil {
		return "Etc/GMT+12", err
	}
	if country.Timezones == nil || len(country.Timezones) == 0 {
		return "Etc/GMT+12", nil
	}
	countryTimezone := country.Timezones[0]
	timezone, ok := slice.FindBy(s.List(), func(index int, item mResource.Timezone) bool {
		return strings.Contains(item.Description, countryTimezone)
	})
	if !ok {
		return "Etc/GMT+12", nil
	}
	return timezone.OlsonName, nil
}
