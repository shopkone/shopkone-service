package sDomain

import (
	"shopkone-service/internal/module/setting/domains/mDomains"
	"shopkone-service/utility/code"
)

func (s *sDomain) PreCheck(domain string) (data DomainBindInfoOut, err error) {
	// 判断域名是否在服务商处已注册
	if err = s.DomainIsRegister(domain); err != nil {
		return data, err
	}

	// 域名是否已被连接
	var count int64
	if err = s.orm.Model(&mDomains.Domain{}).Where("domain = ?", domain).
		Where("status = ?", mDomains.DomainStatusConnectSuccess).
		Count(&count).Error; err != nil {
		return data, err
	}
	if count > 0 {
		return data, code.DomainAlreadyConnected
	}

	// 获取可用的ip
	data, err = s.DomainBindInfo()
	if err != nil {
		return data, err
	}

	return data, err
}
