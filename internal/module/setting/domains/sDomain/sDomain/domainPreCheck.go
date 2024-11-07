package sDomain

import (
	"shopkone-service/internal/module/setting/domains/mDomains"
	"shopkone-service/utility/code"
)

type PreCheckOut struct {
	IP string
}

func (s *sDomain) PreCheck(domain string, IsShopkone bool) (data PreCheckOut, err error) {
	// 判断域名是否已注册
	if err = s.DomainIsRegister(domain); err != nil {
		return PreCheckOut{}, err
	}

	// 域名是否已被连接
	var count int64
	if err := s.orm.Model(&mDomains.Domain{}).Where("domain = ?", domain).
		Where("status = ?", mDomains.DomainStatusConnectPre).
		Count(&count).Error; err != nil {
		return data, err
	}
	if count > 0 {
		return data, code.DomainAlreadyConnected
	}

	// 获取可用的ip
	ip, err := s.GetIP()
	if err != nil {
		return data, err
	}

	// 创建域名连接
	data.IP = ip
	return data, err
}
