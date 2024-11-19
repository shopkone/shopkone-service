package sBlackIp

import (
	"shopkone-service/internal/module/setting/domains/mDomains"
)

func (s *sBlackIp) List(t *mDomains.BlackIpType) (ips []mDomains.DomainBlackIp, err error) {
	query := s.orm.Model(&mDomains.DomainBlackIp{}).
		Where("shop_id = ?", s.shopId)
	if t != nil {
		query = query.Where("type = ?", t)
	}
	query = query.Select("ip", "type")
	err = query.Find(&ips).Error
	if err != nil {
		return ips, err
	}
	return ips, err
}
