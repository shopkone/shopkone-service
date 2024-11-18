package sBlockCountry

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/base/resource/sResource"
	"shopkone-service/internal/module/setting/domains/mDomains"
)

func (s *sBlockCountry) Update(codes []string) (err error) {
	if codes == nil {
		codes = []string{}
	}
	// 校验国家是否存在
	if err = sResource.NewCountry().CheckCountryListExist(codes); err != nil {
		return err
	}
	// 获取旧的列表
	oldCode, err := s.List()
	// 对比两者
	// 删除
	removes := slice.Filter(oldCode, func(index int, item string) bool {
		return !slice.Contain(codes, item)
	})
	if len(removes) > 0 {
		if err = s.orm.Model(&mDomains.DomainBlockCountry{}).
			Where("country_code IN ? AND shop_id = ?", removes, s.shopId).
			Unscoped().
			Delete(&mDomains.DomainBlockCountry{}).Error; err != nil {
			return err
		}
	}
	// 新增
	adds := slice.Filter(codes, func(index int, item string) bool {
		return !slice.Contain(oldCode, item)
	})
	if len(adds) > 0 {
		addList := slice.Map(adds, func(index int, item string) mDomains.DomainBlockCountry {
			i := mDomains.DomainBlockCountry{}
			i.CountryCode = item
			i.ShopId = s.shopId
			return i
		})
		return s.orm.Create(&addList).Error
	}
	return err
}
