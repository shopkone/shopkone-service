package sDomain

import (
	"golang.org/x/net/publicsuffix"
	"shopkone-service/utility/code"
)

func (s *sDomain) DomainIsRegister(domain string) error {
	// 判断是否是子域名
	if true {
		return code.DomainNotRegistered
	}
	publicsuffix.PublicSuffix(domain)
	return nil
}
