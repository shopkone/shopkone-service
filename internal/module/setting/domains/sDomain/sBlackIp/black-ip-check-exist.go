package sBlackIp

import (
	"shopkone-service/internal/module/setting/domains/mDomains"
	"shopkone-service/utility/code"
)

func (s *sBlackIp) CheckExist(ips []string, t mDomains.BlackIpType) (err error) {
	var count int64
	err = s.orm.Model(&mDomains.DomainBlackIp{}).
		Where("shop_id = ?", s.shopId).
		Where("ip IN (?) AND type = ?", ips, t).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return code.DomainBlackIpExist
	}
	return err
}
