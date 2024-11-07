package sDomain

import (
	"golang.org/x/net/publicsuffix"
	"shopkone-service/internal/module/base/ali/sAli"
	"shopkone-service/utility/code"
)

func (s *sDomain) DomainIsRegister(domain string) error {
	mainDomain, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return err
	}
	isRegister, err := sAli.DomainIsRegister(mainDomain)
	if err != nil {
		return err
	}
	// 如果未注册，则报错
	if !isRegister {
		return code.DomainNotRegistered
	}
	return nil
}
