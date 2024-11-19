package sBlackIp

import (
	"github.com/duke-git/lancet/v2/slice"
	"net"
	"shopkone-service/internal/module/setting/domains/mDomains"
	"shopkone-service/utility/code"
)

func (s *sBlackIp) Update(ips []string, t mDomains.BlackIpType) (err error) {
	// 校验类型
	if t != mDomains.BlackIpTypeBlack && t != mDomains.BlackIpTypeWhite {
		return code.DomainValidBlackIpType
	}

	// 校验ip可用性
	isAllValid := slice.Every(ips, func(index int, ip string) bool {
		if net.ParseIP(ip) == nil {
			return false
		}
		return true
	})
	if !isAllValid {
		return code.DomainValidIP
	}

	// 判断是否存在于该店铺
	if err = s.CheckExist(ips, t); err != nil {
		return err
	}

	// 添加
	if err = s.Add(ips, t); err != nil {
		return err
	}

	// 移除
	list, err := s.List(&t)
	if err != nil {
		return err
	}
	otherType := mDomains.BlackIpTypeBlack
	if t == mDomains.BlackIpTypeBlack {
		otherType = mDomains.BlackIpTypeWhite
	} else if t == mDomains.BlackIpTypeWhite {
		otherType = mDomains.BlackIpTypeBlack
	}
	listIps := slice.Map(list, func(index int, item mDomains.DomainBlackIp) string {
		return item.Ip
	})
	return s.Remove(listIps, otherType)
}
