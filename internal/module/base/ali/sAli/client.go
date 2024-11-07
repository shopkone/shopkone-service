package sAli

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	domain20180129 "github.com/alibabacloud-go/domain-20180129/v5/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"shopkone-service/hack"
)

var stsClient *sts20150401.Client
var domainClient *domain20180129.Client

func AliYunClient() error {
	// 获取配置
	conf, err := hack.GetConfig()
	if err != nil {
		return err
	}
	config := &openapi.Config{
		AccessKeyId:     &conf.AliYun.AccessKeyId,
		AccessKeySecret: &conf.AliYun.AccessKeySecret,
		Endpoint:        &conf.AliYun.EndPoint,
	}
	// 初始化客户端
	stsClient, err = sts20150401.NewClient(config)
	if err != nil {
		return err
	}
	// 初始化domain客户端
	domainConfig := &openapi.Config{
		AccessKeyId:     tea.String("LTAI5tSLPHLVTscYd2fejkPi"),
		AccessKeySecret: tea.String("0Jz95TivfRi8e8lBCvCLwVepoC8VIr"),
		Endpoint:        tea.String("domain.aliyuncs.com"),
	}
	domainClient, err = domain20180129.NewClient(domainConfig)
	return err
}
