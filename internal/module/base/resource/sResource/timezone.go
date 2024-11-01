package sResource

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/base/resource"
	"shopkone-service/internal/module/base/resource/mResource"
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
	data := slice.Filter(resource.CountryTimezones, func(index int, item mResource.CountryTimeZone) bool {
		return item.CountryCode == c
	})
	if len(data) == 0 {
		return "Etc/GMT+12", nil
	}
	find, ok := slice.FindBy(data, func(index int, item mResource.CountryTimeZone) bool {
		_, o := slice.FindBy(resource.Timezones, func(index int, i mResource.Timezone) bool {
			return item.Timezone == i.OlsonName
		})
		return o
	})
	if !ok {
		return "Etc/GMT+12", nil
	}
	return find.Timezone, nil
}
