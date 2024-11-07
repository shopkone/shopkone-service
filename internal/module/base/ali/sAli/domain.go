package sAli

import (
	domain20180129 "github.com/alibabacloud-go/domain-20180129/v5/client"
	"github.com/alibabacloud-go/tea/tea"
	"shopkone-service/utility/code"
)

func DomainIsRegister(domain string) (is bool, err error) {
	checkDomainRequest := &domain20180129.CheckDomainRequest{
		DomainName: tea.String(domain),
	}
	ret, err := domainClient.CheckDomain(checkDomainRequest)
	if err != nil {
		return is, err
	}
	// 可以注册，说明未注册
	if *ret.Body.Avail == "1" {
		return false, err
	}
	// -1异常，直接报错：域名异常
	if *ret.Body.Avail == "-1" {
		return false, code.DomainUnknown
	}
	// 其他情况，都默认为已注册
	return true, err
}
