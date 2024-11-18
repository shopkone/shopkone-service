package sBlockCountry

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/domains/mDomains"
)

func (s *sBlockCountry) List() (codes []string, err error) {
	var list []mDomains.DomainBlockCountry
	err = s.orm.Model(&mDomains.DomainBlockCountry{}).
		Where("shop_id = ?", s.shopId).
		Select("country_code").
		Find(&list).
		Error
	if err != nil {
		return nil, err
	}
	codes = slice.Map(list, func(index int, item mDomains.DomainBlockCountry) string {
		return item.CountryCode
	})
	if codes == nil {
		codes = []string{}
	}
	return codes, err
}
