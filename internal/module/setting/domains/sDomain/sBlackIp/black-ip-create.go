package sBlackIp

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/module/setting/domains/mDomains"
)

func (s *sBlackIp) Add(ips []string, t mDomains.BlackIpType) (err error) {
	if len(ips) == 0 {
		return err
	}
	ipList := slice.Map(ips, func(index int, item string) mDomains.DomainBlackIp {
		i := mDomains.DomainBlackIp{
			Ip:   item,
			Type: t,
		}
		i.ShopId = s.shopId
		return i
	})
	return s.orm.Create(&ipList).Error
}
