package sBlackIp

import (
	"shopkone-service/internal/module/setting/domains/mDomains"
)

func (s *sBlackIp) Remove(ips []string, t mDomains.BlackIpType) (err error) {
	if len(ips) == 0 {
		return err
	}
	err = s.orm.Model(&mDomains.DomainBlackIp{}).
		Unscoped().
		Where("shop_id = ?", s.shopId).
		Where("ip IN (?) AND type = ?", ips, t).
		Delete(&mDomains.DomainBlackIp{}).Error
	return err
}
